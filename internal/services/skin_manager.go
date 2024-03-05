package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const (
	url               = "https://bymykel.github.io/CSGO-API/api/ru/crates/cases.json"
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
	cases []Case
}

func NewSkinManager() *SkinManager {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error executing the GET request:", err)
	}
	defer response.Body.Close()

	var cases []Case
	if err := json.NewDecoder(response.Body).Decode(&cases); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	return &SkinManager{cases: cases}
}

type Case struct {
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
	Case    string
	Pattern string
	Float   string
}

type CaseDto struct {
	Name string
}

func (sm *SkinManager) GetSkin(partCaseName string) (*SkinDto, error) {
	caseId, err := sm.findCaseIdByPartName(partCaseName)
	if err != nil {
		return nil, err
	}

	skinType, isRare := getSkinType()

	skin, err := sm.getSkin(caseId, skinType, isRare)
	if err != nil {
		return nil, err
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
		Case:    sm.cases[caseId].Name,
		Pattern: strconv.Itoa(rand.Int() % 1001),
		Float:   strconv.FormatFloat(getFloat(skin.Rarity.ID), 'f', 3, 64),
	}

	return &resultSkin, nil
}

func (sm *SkinManager) GetCases() []CaseDto {
	var cases []CaseDto
	for _, c := range sm.cases {
		cases = append(cases, CaseDto{Name: c.Name})
	}

	return cases
}

func (sm *SkinManager) findCaseIdByPartName(partName string) (int, error) {
	maxMatchId := -1
	maxMatchCount := 0

	for id, c := range sm.cases {
		matchCount := strings.Count(normalize(c.Name), normalize(partName))
		if matchCount > maxMatchCount {
			maxMatchId = id
			maxMatchCount = matchCount
		}
	}

	if maxMatchId == -1 {
		return 0, errors.New("Кейс не найден")
	}

	return maxMatchId, nil
}

func (sm *SkinManager) getSkin(caseId int, skinType string, isRare bool) (*Skin, error) {
	c := sm.cases[caseId]
	if isRare {
		count := len(c.Rares)
		if count == 0 {
			return nil, errors.New("Нулевое количество ножей у кейса с id=" + strconv.Itoa(caseId))
		}
		randNum := rand.Int() % count
		return &c.Rares[randNum], nil
	}

	skins := sm.getSkins(caseId, skinType)
	count := len(skins)
	if count == 0 {
		return nil, errors.New("Нулевое количество скинов типа " + skinType + " у кейса с id=" + strconv.Itoa(caseId))
	}

	randNum := rand.Int() % count
	return &skins[randNum], nil
}

func (sm *SkinManager) getSkins(caseId int, skinType string) []Skin {
	var skins []Skin
	for _, s := range sm.cases[caseId].Weapons {
		if s.Rarity.ID == skinType {
			skins = append(skins, s)
		}
	}

	return skins
}

func getSkinType() (string, bool) {
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
	default:
		return rareName
	}
}

func normalize(str string) string {
	return strings.ToLower(strings.ReplaceAll(str, "ё", "е"))
}

func getFloat(skinType string) float64 {
	if skinType == ancientGloveId {
		return float64(rand.Int()%940+60) / 1000
	}

	return float64(rand.Int()%1000) / 1000
}
