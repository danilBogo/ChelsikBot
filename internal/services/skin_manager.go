package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

const (
	casesUrl          = "https://bymykel.github.io/CSGO-API/api/ru/crates/cases.json"
	collectionsUrl    = "https://bymykel.github.io/CSGO-API/api/ru/collections.json"
	commonId          = "rarity_common_weapon"
	commonName        = "\U0001FA76\U0001FA76\U0001FA76\U0001FA76\U0001FA76\U0001FA76\U0001FA76\U0001FA76\U0001FA76\U0001FA76"
	uncommonId        = "rarity_uncommon_weapon"
	uncommonName      = "\U0001FA75\U0001FA75\U0001FA75\U0001FA75\U0001FA75\U0001FA75\U0001FA75\U0001FA75\U0001FA75\U0001FA75"
	rareId            = "rarity_rare_weapon"
	rareName          = "💙💙💙💙💙💙💙💙💙💙"
	mythicalId        = "rarity_mythical_weapon"
	mythicalName      = "💜💜💜💜💜💜💜💜💜💜"
	legendaryId       = "rarity_legendary_weapon"
	legendaryName     = "\U0001FA77\U0001FA77\U0001FA77\U0001FA77\U0001FA77\U0001FA77\U0001FA77\U0001FA77\U0001FA77\U0001FA77"
	ancientWeaponId   = "rarity_ancient_weapon"
	ancientNameWeapon = "❤️❤️❤️❤️❤️❤️❤️❤️❤️❤️"
	ancientNameKnife  = "💛💛💛💛💛💛💛💛💛💛"
	ancientGloveId    = "rarity_ancient"
	ancientNameGlove  = "💛💛💛💛💛💛💛💛💛💛"
)

type SkinManager struct {
	sets            []Set
	CasesName       []string
	CollectionsName []string
}

func NewSkinManager() *SkinManager {
	casesResponse, err := http.Get(casesUrl)
	if err != nil {
		log.Fatal("Error executing the GET request:", err)
	}
	defer casesResponse.Body.Close()

	var cases []Set
	if err := json.NewDecoder(casesResponse.Body).Decode(&cases); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	casesNames := make([]string, len(cases))
	for id, c := range cases {
		casesNames[id] = c.Name
	}

	collectionsResponse, err := http.Get(collectionsUrl)
	if err != nil {
		log.Fatal("Error executing the GET request:", err)
	}
	defer collectionsResponse.Body.Close()

	var collections []Set
	if err := json.NewDecoder(collectionsResponse.Body).Decode(&collections); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	var collectionsWithoutAgentsAndGraffiti []Set
	for _, c := range collections {
		if strings.Contains(strings.ToLower(c.Name), "коллекция") &&
			!strings.Contains(strings.ToLower(c.Name), "граффити") {
			collectionsWithoutAgentsAndGraffiti = append(collectionsWithoutAgentsAndGraffiti, c)
		}
	}

	collectionsNames := make([]string, len(collectionsWithoutAgentsAndGraffiti))
	for id, c := range collectionsWithoutAgentsAndGraffiti {
		collectionsNames[id] = c.Name
	}

	sets := slices.Concat(cases, collectionsWithoutAgentsAndGraffiti)

	return &SkinManager{sets: sets, CasesName: casesNames, CollectionsName: collectionsNames}
}

type Set struct {
	Name    string `json:"name"`
	Weapons []Skin `json:"contains"`
	Rares   []Skin `json:"contains_rare"`
	Image   string `json:"image"`
}

type Skin struct {
	Name   string `json:"name"`
	Rarity Rarity `json:"rarity"`
	Phase  any    `json:"phase,omitempty"`
	Image  string `json:"image"`
}

type Rarity struct {
	ID string `json:"id"`
}

type SkinDto struct {
	Name    string
	Rarity  string
	Phase   any
	Image   []byte
	Set     string
	Pattern string
	Float   string
}

func (sm *SkinManager) GetSkin(partCaseName string) (*SkinDto, error) {
	setId, err := sm.findSetIdByPartName(partCaseName)
	if err != nil {
		return nil, err
	}

	var skinType string
	var isRare bool
	var skin *Skin
	if strings.Contains(strings.ToLower(sm.sets[setId].Name), "коллекция") {
		skinType = getCollectionSkinType()
		skin = sm.getCollectionSkin(setId, skinType, isRare)
	} else {
		skinType, isRare = getCaseSkinType()
		skin, err = sm.getSkin(setId, skinType, isRare)
		if err != nil {
			return nil, err
		}
	}

	resp, err := http.Get(skin.Image)
	if err != nil {
		log.Fatal("Error executing the GET request:", err)
	}
	defer resp.Body.Close()

	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading file contents:", err)
	}

	resultSkin := SkinDto{
		Name:    skin.Name,
		Rarity:  convertSkinTypeToName(skin.Rarity.ID, isRare),
		Phase:   skin.Phase,
		Image:   fileBytes,
		Set:     sm.sets[setId].Name,
		Pattern: strconv.Itoa(rand.Int() % 1001),
		Float:   strconv.FormatFloat(getFloat(skin.Rarity.ID), 'f', 3, 64),
	}

	return &resultSkin, nil
}

func (sm *SkinManager) findSetIdByPartName(partName string) (int, error) {
	maxMatchId := -1
	maxMatchCount := 0

	for id, c := range sm.sets {
		matchCount := strings.Count(normalize(c.Name), normalize(partName))
		if matchCount > maxMatchCount {
			maxMatchId = id
			maxMatchCount = matchCount
		}
	}

	if maxMatchId == -1 {
		return 0, errors.New("Кейс или коллекция не найдены")
	}

	return maxMatchId, nil
}

func (sm *SkinManager) getCollectionSkin(setId int, skinType string, isRare bool) *Skin {
	currentSkinType := skinType
	for {
		skin, err := sm.getSkin(setId, currentSkinType, isRare)
		if err == nil {
			return skin
		}

		currentSkinType = getCollectionSkinType()
	}
}

func (sm *SkinManager) getSkin(setId int, skinType string, isRare bool) (*Skin, error) {
	c := sm.sets[setId]
	if isRare {
		count := len(c.Rares)
		if count == 0 {
			return nil, errors.New("Нулевое количество ножей у кейса с id=" + strconv.Itoa(setId))
		}
		randNum := rand.Int() % count
		return &c.Rares[randNum], nil
	}

	skins := sm.getSkins(setId, skinType)
	count := len(skins)
	if count == 0 {
		return nil, errors.New("Нулевое количество скинов типа " + skinType + " у кейса с id=" + strconv.Itoa(setId))
	}

	randNum := rand.Int() % count
	return &skins[randNum], nil
}

func (sm *SkinManager) getSkins(setId int, skinType string) []Skin {
	var skins []Skin
	for _, s := range sm.sets[setId].Weapons {
		if s.Rarity.ID == skinType {
			skins = append(skins, s)
		}
	}

	return skins
}

func getCaseSkinType() (string, bool) {
	randNum := rand.Int() % 10000
	switch {
	case randNum >= 0 && randNum < 26:
		return ancientWeaponId, true
	case randNum >= 26 && randNum < 90:
		return ancientWeaponId, false
	case randNum >= 90 && randNum < 400:
		return legendaryId, false
	case randNum >= 400 && randNum < 2000:
		return mythicalId, false
	default:
		return rareId, false
	}
}

func getCollectionSkinType() string {
	randNum := rand.Int() % 10000
	switch {
	case randNum >= 0 && randNum < 5:
		return ancientWeaponId
	case randNum >= 5 && randNum < 26:
		return legendaryId
	case randNum >= 26 && randNum < 124:
		return mythicalId
	case randNum >= 124 && randNum < 545:
		return rareId
	case randNum >= 545 && randNum < 2353:
		return uncommonId
	default:
		return commonId
	}
}

func convertSkinTypeToName(skinType string, isRare bool) string {
	switch {
	case skinType == ancientWeaponId && isRare:
		return ancientNameKnife
	case skinType == ancientWeaponId:
		return ancientNameWeapon
	case skinType == ancientGloveId:
		return ancientNameGlove
	case skinType == legendaryId:
		return legendaryName
	case skinType == mythicalId:
		return mythicalName
	case skinType == rareId:
		return rareName
	case skinType == uncommonId:
		return uncommonName
	default:
		return commonName
	}
}

func normalize(str string) string {
	remove := "«»\"'"

	filter := func(r rune) rune {
		if strings.ContainsRune(remove, r) {
			return -1
		}
		return r
	}

	result := strings.Map(filter, str)
	return strings.ToLower(strings.ReplaceAll(result, "ё", "е"))
}

func getFloat(skinType string) float64 {
	if skinType == ancientGloveId {
		return float64(rand.Int()%940+60) / 1000
	}

	return float64(rand.Int()%1000) / 1000
}
