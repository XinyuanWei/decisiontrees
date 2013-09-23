// Code generated by protoc-gen-go.
// source: protobufs/decisiontrees.proto
// DO NOT EDIT!

package protobufs

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type LossFunction int32

const (
	LossFunction_LOGIT                    LossFunction = 1
	LossFunction_LEAST_ABSOLUTE_DEVIATION LossFunction = 2
	LossFunction_HUBER                    LossFunction = 3
)

var LossFunction_name = map[int32]string{
	1: "LOGIT",
	2: "LEAST_ABSOLUTE_DEVIATION",
	3: "HUBER",
}
var LossFunction_value = map[string]int32{
	"LOGIT":                    1,
	"LEAST_ABSOLUTE_DEVIATION": 2,
	"HUBER":                    3,
}

func (x LossFunction) Enum() *LossFunction {
	p := new(LossFunction)
	*p = x
	return p
}
func (x LossFunction) String() string {
	return proto.EnumName(LossFunction_name, int32(x))
}
func (x LossFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LossFunction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LossFunction_value, data, "LossFunction")
	if err != nil {
		return err
	}
	*x = LossFunction(value)
	return nil
}

type Algorithm int32

const (
	Algorithm_BOOSTING      Algorithm = 1
	Algorithm_RANDOM_FOREST Algorithm = 2
)

var Algorithm_name = map[int32]string{
	1: "BOOSTING",
	2: "RANDOM_FOREST",
}
var Algorithm_value = map[string]int32{
	"BOOSTING":      1,
	"RANDOM_FOREST": 2,
}

func (x Algorithm) Enum() *Algorithm {
	p := new(Algorithm)
	*p = x
	return p
}
func (x Algorithm) String() string {
	return proto.EnumName(Algorithm_name, int32(x))
}
func (x Algorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Algorithm) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Algorithm_value, data, "Algorithm")
	if err != nil {
		return err
	}
	*x = Algorithm(value)
	return nil
}

type TrainingStatus int32

const (
	TrainingStatus_UNCLAIMED  TrainingStatus = 1
	TrainingStatus_PROCESSING TrainingStatus = 2
	TrainingStatus_FINISHED   TrainingStatus = 3
)

var TrainingStatus_name = map[int32]string{
	1: "UNCLAIMED",
	2: "PROCESSING",
	3: "FINISHED",
}
var TrainingStatus_value = map[string]int32{
	"UNCLAIMED":  1,
	"PROCESSING": 2,
	"FINISHED":   3,
}

func (x TrainingStatus) Enum() *TrainingStatus {
	p := new(TrainingStatus)
	*p = x
	return p
}
func (x TrainingStatus) String() string {
	return proto.EnumName(TrainingStatus_name, int32(x))
}
func (x TrainingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *TrainingStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TrainingStatus_value, data, "TrainingStatus")
	if err != nil {
		return err
	}
	*x = TrainingStatus(value)
	return nil
}

type DataSource int32

const (
	DataSource_GRIDFS DataSource = 1
)

var DataSource_name = map[int32]string{
	1: "GRIDFS",
}
var DataSource_value = map[string]int32{
	"GRIDFS": 1,
}

func (x DataSource) Enum() *DataSource {
	p := new(DataSource)
	*p = x
	return p
}
func (x DataSource) String() string {
	return proto.EnumName(DataSource_name, int32(x))
}
func (x DataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DataSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataSource_value, data, "DataSource")
	if err != nil {
		return err
	}
	*x = DataSource(value)
	return nil
}

type Feature struct {
	Feature          *int64   `protobuf:"varint,1,opt,name=feature" json:"feature,omitempty" bson:"feature,omitempty"`
	Value            *float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty" bson:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-" bson:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}

func (m *Feature) GetFeature() int64 {
	if m != nil && m.Feature != nil {
		return *m.Feature
	}
	return 0
}

func (m *Feature) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Example struct {
	Label            *float64  `protobuf:"fixed64,1,opt,name=label" json:"label,omitempty" bson:"label,omitempty"`
	WeightedLabel    *float64  `protobuf:"fixed64,2,opt,name=weightedLabel" json:"weightedLabel,omitempty" bson:"weightedLabel,omitempty"`
	Features         []float64 `protobuf:"fixed64,3,rep,packed,name=features" json:"features,omitempty" bson:"features,omitempty"`
	XXX_unrecognized []byte    `json:"-" bson:"-"`
}

func (m *Example) Reset()         { *m = Example{} }
func (m *Example) String() string { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()    {}

func (m *Example) GetLabel() float64 {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return 0
}

func (m *Example) GetWeightedLabel() float64 {
	if m != nil && m.WeightedLabel != nil {
		return *m.WeightedLabel
	}
	return 0
}

func (m *Example) GetFeatures() []float64 {
	if m != nil {
		return m.Features
	}
	return nil
}

type TrainingData struct {
	Train            []*Example `protobuf:"bytes,1,rep,name=train" json:"train,omitempty" bson:"train,omitempty"`
	Test             []*Example `protobuf:"bytes,2,rep,name=test" json:"test,omitempty" bson:"test,omitempty"`
	XXX_unrecognized []byte     `json:"-" bson:"-"`
}

func (m *TrainingData) Reset()         { *m = TrainingData{} }
func (m *TrainingData) String() string { return proto.CompactTextString(m) }
func (*TrainingData) ProtoMessage()    {}

func (m *TrainingData) GetTrain() []*Example {
	if m != nil {
		return m.Train
	}
	return nil
}

func (m *TrainingData) GetTest() []*Example {
	if m != nil {
		return m.Test
	}
	return nil
}

type TreeNode struct {
	// feature to split on
	Feature *int64 `protobuf:"varint,1,opt,name=feature" json:"feature,omitempty" bson:"feature,omitempty"`
	// value to split on
	SplitValue       *float64    `protobuf:"fixed64,2,opt,name=splitValue" json:"splitValue,omitempty" bson:"splitValue,omitempty"`
	Left             *TreeNode   `protobuf:"bytes,3,opt,name=left" json:"left,omitempty" bson:"left,omitempty"`
	Right            *TreeNode   `protobuf:"bytes,4,opt,name=right" json:"right,omitempty" bson:"right,omitempty"`
	LeafValue        *float64    `protobuf:"fixed64,5,opt,name=leafValue" json:"leafValue,omitempty" bson:"leafValue,omitempty"`
	Annotation       *Annotation `protobuf:"bytes,6,opt,name=annotation" json:"annotation,omitempty" bson:"annotation,omitempty"`
	XXX_unrecognized []byte      `json:"-" bson:"-"`
}

func (m *TreeNode) Reset()         { *m = TreeNode{} }
func (m *TreeNode) String() string { return proto.CompactTextString(m) }
func (*TreeNode) ProtoMessage()    {}

func (m *TreeNode) GetFeature() int64 {
	if m != nil && m.Feature != nil {
		return *m.Feature
	}
	return 0
}

func (m *TreeNode) GetSplitValue() float64 {
	if m != nil && m.SplitValue != nil {
		return *m.SplitValue
	}
	return 0
}

func (m *TreeNode) GetLeft() *TreeNode {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *TreeNode) GetRight() *TreeNode {
	if m != nil {
		return m.Right
	}
	return nil
}

func (m *TreeNode) GetLeafValue() float64 {
	if m != nil && m.LeafValue != nil {
		return *m.LeafValue
	}
	return 0
}

func (m *TreeNode) GetAnnotation() *Annotation {
	if m != nil {
		return m.Annotation
	}
	return nil
}

type Annotation struct {
	NumExamples      *int64   `protobuf:"varint,1,opt,name=numExamples" json:"numExamples,omitempty" bson:"numExamples,omitempty"`
	AverageGain      *float64 `protobuf:"fixed64,2,opt,name=averageGain" json:"averageGain,omitempty" bson:"averageGain,omitempty"`
	XXX_unrecognized []byte   `json:"-" bson:"-"`
}

func (m *Annotation) Reset()         { *m = Annotation{} }
func (m *Annotation) String() string { return proto.CompactTextString(m) }
func (*Annotation) ProtoMessage()    {}

func (m *Annotation) GetNumExamples() int64 {
	if m != nil && m.NumExamples != nil {
		return *m.NumExamples
	}
	return 0
}

func (m *Annotation) GetAverageGain() float64 {
	if m != nil && m.AverageGain != nil {
		return *m.AverageGain
	}
	return 0
}

type Forest struct {
	Trees            []*TreeNode `protobuf:"bytes,1,rep,name=trees" json:"trees,omitempty" bson:"trees,omitempty"`
	XXX_unrecognized []byte      `json:"-" bson:"-"`
}

func (m *Forest) Reset()         { *m = Forest{} }
func (m *Forest) String() string { return proto.CompactTextString(m) }
func (*Forest) ProtoMessage()    {}

func (m *Forest) GetTrees() []*TreeNode {
	if m != nil {
		return m.Trees
	}
	return nil
}

type SplittingConstraints struct {
	MaximumLevels        *int64   `protobuf:"varint,1,opt,name=maximumLevels" json:"maximumLevels,omitempty" bson:"maximumLevels,omitempty"`
	MinimumAverageGain   *float64 `protobuf:"fixed64,2,opt,name=minimumAverageGain" json:"minimumAverageGain,omitempty" bson:"minimumAverageGain,omitempty"`
	MinimumSamplesAtLeaf *int64   `protobuf:"varint,3,opt,name=minimumSamplesAtLeaf" json:"minimumSamplesAtLeaf,omitempty" bson:"minimumSamplesAtLeaf,omitempty"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
}

func (m *SplittingConstraints) Reset()         { *m = SplittingConstraints{} }
func (m *SplittingConstraints) String() string { return proto.CompactTextString(m) }
func (*SplittingConstraints) ProtoMessage()    {}

func (m *SplittingConstraints) GetMaximumLevels() int64 {
	if m != nil && m.MaximumLevels != nil {
		return *m.MaximumLevels
	}
	return 0
}

func (m *SplittingConstraints) GetMinimumAverageGain() float64 {
	if m != nil && m.MinimumAverageGain != nil {
		return *m.MinimumAverageGain
	}
	return 0
}

func (m *SplittingConstraints) GetMinimumSamplesAtLeaf() int64 {
	if m != nil && m.MinimumSamplesAtLeaf != nil {
		return *m.MinimumSamplesAtLeaf
	}
	return 0
}

type PruningConstraints struct {
	CrossValidationFolds *int64 `protobuf:"varint,1,opt,name=crossValidationFolds" json:"crossValidationFolds,omitempty" bson:"crossValidationFolds,omitempty"`
	XXX_unrecognized     []byte `json:"-" bson:"-"`
}

func (m *PruningConstraints) Reset()         { *m = PruningConstraints{} }
func (m *PruningConstraints) String() string { return proto.CompactTextString(m) }
func (*PruningConstraints) ProtoMessage()    {}

func (m *PruningConstraints) GetCrossValidationFolds() int64 {
	if m != nil && m.CrossValidationFolds != nil {
		return *m.CrossValidationFolds
	}
	return 0
}

type InfluenceTrimmingConfig struct {
	Alpha            *float64 `protobuf:"fixed64,1,opt,name=alpha" json:"alpha,omitempty" bson:"alpha,omitempty"`
	WarmupRounds     *int64   `protobuf:"varint,2,opt,name=warmupRounds" json:"warmupRounds,omitempty" bson:"warmupRounds,omitempty"`
	XXX_unrecognized []byte   `json:"-" bson:"-"`
}

func (m *InfluenceTrimmingConfig) Reset()         { *m = InfluenceTrimmingConfig{} }
func (m *InfluenceTrimmingConfig) String() string { return proto.CompactTextString(m) }
func (*InfluenceTrimmingConfig) ProtoMessage()    {}

func (m *InfluenceTrimmingConfig) GetAlpha() float64 {
	if m != nil && m.Alpha != nil {
		return *m.Alpha
	}
	return 0
}

func (m *InfluenceTrimmingConfig) GetWarmupRounds() int64 {
	if m != nil && m.WarmupRounds != nil {
		return *m.WarmupRounds
	}
	return 0
}

type LossFunctionConfig struct {
	LossFunction     *LossFunction `protobuf:"varint,1,opt,name=lossFunction,enum=protobufs.LossFunction" json:"lossFunction,omitempty" bson:"lossFunction,omitempty"`
	HuberAlpha       *float64      `protobuf:"fixed64,2,opt,name=huberAlpha" json:"huberAlpha,omitempty" bson:"huberAlpha,omitempty"`
	XXX_unrecognized []byte        `json:"-" bson:"-"`
}

func (m *LossFunctionConfig) Reset()         { *m = LossFunctionConfig{} }
func (m *LossFunctionConfig) String() string { return proto.CompactTextString(m) }
func (*LossFunctionConfig) ProtoMessage()    {}

func (m *LossFunctionConfig) GetLossFunction() LossFunction {
	if m != nil && m.LossFunction != nil {
		return *m.LossFunction
	}
	return 0
}

func (m *LossFunctionConfig) GetHuberAlpha() float64 {
	if m != nil && m.HuberAlpha != nil {
		return *m.HuberAlpha
	}
	return 0
}

type ShrinkageConfig struct {
	Shrinkage        *float64 `protobuf:"fixed64,1,opt,name=shrinkage" json:"shrinkage,omitempty" bson:"shrinkage,omitempty"`
	XXX_unrecognized []byte   `json:"-" bson:"-"`
}

func (m *ShrinkageConfig) Reset()         { *m = ShrinkageConfig{} }
func (m *ShrinkageConfig) String() string { return proto.CompactTextString(m) }
func (*ShrinkageConfig) ProtoMessage()    {}

func (m *ShrinkageConfig) GetShrinkage() float64 {
	if m != nil && m.Shrinkage != nil {
		return *m.Shrinkage
	}
	return 0
}

type StochasticityConfig struct {
	// Take a random sample of training data per round
	// Used in stochastic gradient boosting
	PerRoundSamplingRate *float64 `protobuf:"fixed64,1,opt,name=perRoundSamplingRate" json:"perRoundSamplingRate,omitempty" bson:"perRoundSamplingRate,omitempty"`
	// Proportion of examples to draw boostrap sample from
	// Used in random forests
	ExampleBoostrapProportion *float64 `protobuf:"fixed64,2,opt,name=exampleBoostrapProportion" json:"exampleBoostrapProportion,omitempty" bson:"exampleBoostrapProportion,omitempty"`
	// Number of features to examine at each splitting step
	// Used in random forests
	FeatureSampleSize *int64 `protobuf:"varint,3,opt,name=featureSampleSize" json:"featureSampleSize,omitempty" bson:"featureSampleSize,omitempty"`
	XXX_unrecognized  []byte `json:"-" bson:"-"`
}

func (m *StochasticityConfig) Reset()         { *m = StochasticityConfig{} }
func (m *StochasticityConfig) String() string { return proto.CompactTextString(m) }
func (*StochasticityConfig) ProtoMessage()    {}

func (m *StochasticityConfig) GetPerRoundSamplingRate() float64 {
	if m != nil && m.PerRoundSamplingRate != nil {
		return *m.PerRoundSamplingRate
	}
	return 0
}

func (m *StochasticityConfig) GetExampleBoostrapProportion() float64 {
	if m != nil && m.ExampleBoostrapProportion != nil {
		return *m.ExampleBoostrapProportion
	}
	return 0
}

func (m *StochasticityConfig) GetFeatureSampleSize() int64 {
	if m != nil && m.FeatureSampleSize != nil {
		return *m.FeatureSampleSize
	}
	return 0
}

type ForestConfig struct {
	NumWeakLearners         *int64                   `protobuf:"varint,1,opt,name=numWeakLearners" json:"numWeakLearners,omitempty" bson:"numWeakLearners,omitempty"`
	SplittingConstraints    *SplittingConstraints    `protobuf:"bytes,2,opt,name=splittingConstraints" json:"splittingConstraints,omitempty" bson:"splittingConstraints,omitempty"`
	LossFunctionConfig      *LossFunctionConfig      `protobuf:"bytes,3,opt,name=lossFunctionConfig" json:"lossFunctionConfig,omitempty" bson:"lossFunctionConfig,omitempty"`
	InfluenceTrimmingConfig *InfluenceTrimmingConfig `protobuf:"bytes,4,opt,name=influenceTrimmingConfig" json:"influenceTrimmingConfig,omitempty" bson:"influenceTrimmingConfig,omitempty"`
	ShrinkageConfig         *ShrinkageConfig         `protobuf:"bytes,5,opt,name=shrinkageConfig" json:"shrinkageConfig,omitempty" bson:"shrinkageConfig,omitempty"`
	StochasticityConfig     *StochasticityConfig     `protobuf:"bytes,6,opt,name=stochasticityConfig" json:"stochasticityConfig,omitempty" bson:"stochasticityConfig,omitempty"`
	Algorithm               *Algorithm               `protobuf:"varint,7,opt,name=algorithm,enum=protobufs.Algorithm" json:"algorithm,omitempty" bson:"algorithm,omitempty"`
	XXX_unrecognized        []byte                   `json:"-" bson:"-"`
}

func (m *ForestConfig) Reset()         { *m = ForestConfig{} }
func (m *ForestConfig) String() string { return proto.CompactTextString(m) }
func (*ForestConfig) ProtoMessage()    {}

func (m *ForestConfig) GetNumWeakLearners() int64 {
	if m != nil && m.NumWeakLearners != nil {
		return *m.NumWeakLearners
	}
	return 0
}

func (m *ForestConfig) GetSplittingConstraints() *SplittingConstraints {
	if m != nil {
		return m.SplittingConstraints
	}
	return nil
}

func (m *ForestConfig) GetLossFunctionConfig() *LossFunctionConfig {
	if m != nil {
		return m.LossFunctionConfig
	}
	return nil
}

func (m *ForestConfig) GetInfluenceTrimmingConfig() *InfluenceTrimmingConfig {
	if m != nil {
		return m.InfluenceTrimmingConfig
	}
	return nil
}

func (m *ForestConfig) GetShrinkageConfig() *ShrinkageConfig {
	if m != nil {
		return m.ShrinkageConfig
	}
	return nil
}

func (m *ForestConfig) GetStochasticityConfig() *StochasticityConfig {
	if m != nil {
		return m.StochasticityConfig
	}
	return nil
}

func (m *ForestConfig) GetAlgorithm() Algorithm {
	if m != nil && m.Algorithm != nil {
		return *m.Algorithm
	}
	return 0
}

type GridFsConfig struct {
	Database         *string `protobuf:"bytes,1,opt,name=database" json:"database,omitempty" bson:"database,omitempty"`
	Collection       *string `protobuf:"bytes,2,opt,name=collection,def=fs" json:"collection,omitempty" bson:"collection,omitempty"`
	File             *string `protobuf:"bytes,3,opt,name=file" json:"file,omitempty" bson:"file,omitempty"`
	XXX_unrecognized []byte  `json:"-" bson:"-"`
}

func (m *GridFsConfig) Reset()         { *m = GridFsConfig{} }
func (m *GridFsConfig) String() string { return proto.CompactTextString(m) }
func (*GridFsConfig) ProtoMessage()    {}

const Default_GridFsConfig_Collection string = "fs"

func (m *GridFsConfig) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *GridFsConfig) GetCollection() string {
	if m != nil && m.Collection != nil {
		return *m.Collection
	}
	return Default_GridFsConfig_Collection
}

func (m *GridFsConfig) GetFile() string {
	if m != nil && m.File != nil {
		return *m.File
	}
	return ""
}

type DataSourceConfig struct {
	DataSource       *DataSource   `protobuf:"varint,1,opt,name=dataSource,enum=protobufs.DataSource" json:"dataSource,omitempty" bson:"dataSource,omitempty"`
	GridFsConfig     *GridFsConfig `protobuf:"bytes,2,opt,name=gridFsConfig" json:"gridFsConfig,omitempty" bson:"gridFsConfig,omitempty"`
	XXX_unrecognized []byte        `json:"-" bson:"-"`
}

func (m *DataSourceConfig) Reset()         { *m = DataSourceConfig{} }
func (m *DataSourceConfig) String() string { return proto.CompactTextString(m) }
func (*DataSourceConfig) ProtoMessage()    {}

func (m *DataSourceConfig) GetDataSource() DataSource {
	if m != nil && m.DataSource != nil {
		return *m.DataSource
	}
	return 0
}

func (m *DataSourceConfig) GetGridFsConfig() *GridFsConfig {
	if m != nil {
		return m.GridFsConfig
	}
	return nil
}

type EpochResult struct {
	Roc               *float64 `protobuf:"fixed64,1,opt,name=roc" json:"roc,omitempty" bson:"roc,omitempty"`
	LogScore          *float64 `protobuf:"fixed64,2,opt,name=logScore" json:"logScore,omitempty" bson:"logScore,omitempty"`
	NormalizedEntropy *float64 `protobuf:"fixed64,3,opt,name=normalizedEntropy" json:"normalizedEntropy,omitempty" bson:"normalizedEntropy,omitempty"`
	Calibration       *float64 `protobuf:"fixed64,4,opt,name=calibration" json:"calibration,omitempty" bson:"calibration,omitempty"`
	XXX_unrecognized  []byte   `json:"-" bson:"-"`
}

func (m *EpochResult) Reset()         { *m = EpochResult{} }
func (m *EpochResult) String() string { return proto.CompactTextString(m) }
func (*EpochResult) ProtoMessage()    {}

func (m *EpochResult) GetRoc() float64 {
	if m != nil && m.Roc != nil {
		return *m.Roc
	}
	return 0
}

func (m *EpochResult) GetLogScore() float64 {
	if m != nil && m.LogScore != nil {
		return *m.LogScore
	}
	return 0
}

func (m *EpochResult) GetNormalizedEntropy() float64 {
	if m != nil && m.NormalizedEntropy != nil {
		return *m.NormalizedEntropy
	}
	return 0
}

func (m *EpochResult) GetCalibration() float64 {
	if m != nil && m.Calibration != nil {
		return *m.Calibration
	}
	return 0
}

type TrainingResults struct {
	EpochResults     []*EpochResult `protobuf:"bytes,1,rep,name=epochResults" json:"epochResults,omitempty" bson:"epochResults,omitempty"`
	XXX_unrecognized []byte         `json:"-" bson:"-"`
}

func (m *TrainingResults) Reset()         { *m = TrainingResults{} }
func (m *TrainingResults) String() string { return proto.CompactTextString(m) }
func (*TrainingResults) ProtoMessage()    {}

func (m *TrainingResults) GetEpochResults() []*EpochResult {
	if m != nil {
		return m.EpochResults
	}
	return nil
}

type TrainingRow struct {
	ForestConfig     *ForestConfig     `protobuf:"bytes,1,opt,name=forestConfig" json:"forestConfig,omitempty" bson:"forestConfig,omitempty"`
	Forest           *Forest           `protobuf:"bytes,2,opt,name=forest" json:"forest,omitempty" bson:"forest,omitempty"`
	DataSourceConfig *DataSourceConfig `protobuf:"bytes,3,opt,name=dataSourceConfig" json:"dataSourceConfig,omitempty" bson:"dataSourceConfig,omitempty"`
	TrainingStatus   *TrainingStatus   `protobuf:"varint,4,opt,name=trainingStatus,enum=protobufs.TrainingStatus" json:"trainingStatus,omitempty" bson:"trainingStatus,omitempty"`
	TrainingResults  *TrainingResults  `protobuf:"bytes,5,opt,name=trainingResults" json:"trainingResults,omitempty" bson:"trainingResults,omitempty"`
	XXX_unrecognized []byte            `json:"-" bson:"-"`
}

func (m *TrainingRow) Reset()         { *m = TrainingRow{} }
func (m *TrainingRow) String() string { return proto.CompactTextString(m) }
func (*TrainingRow) ProtoMessage()    {}

func (m *TrainingRow) GetForestConfig() *ForestConfig {
	if m != nil {
		return m.ForestConfig
	}
	return nil
}

func (m *TrainingRow) GetForest() *Forest {
	if m != nil {
		return m.Forest
	}
	return nil
}

func (m *TrainingRow) GetDataSourceConfig() *DataSourceConfig {
	if m != nil {
		return m.DataSourceConfig
	}
	return nil
}

func (m *TrainingRow) GetTrainingStatus() TrainingStatus {
	if m != nil && m.TrainingStatus != nil {
		return *m.TrainingStatus
	}
	return 0
}

func (m *TrainingRow) GetTrainingResults() *TrainingResults {
	if m != nil {
		return m.TrainingResults
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobufs.LossFunction", LossFunction_name, LossFunction_value)
	proto.RegisterEnum("protobufs.Algorithm", Algorithm_name, Algorithm_value)
	proto.RegisterEnum("protobufs.TrainingStatus", TrainingStatus_name, TrainingStatus_value)
	proto.RegisterEnum("protobufs.DataSource", DataSource_name, DataSource_value)
}
