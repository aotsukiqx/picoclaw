# Graph Report - .  (2026-04-08)

## Corpus Check
- Large corpus: 965 files · ~1,223,151 words. Semantic extraction will be expensive (many Claude tokens). Consider running on a subfolder, or use --no-semantic to run AST-only.

## Summary
- 7110 nodes · 10739 edges · 232 communities detected
- Extraction: 72% EXTRACTED · 28% INFERRED · 0% AMBIGUOUS · INFERRED: 2987 edges (avg confidence: 0.5)
- Token cost: 0 input · 0 output

## God Nodes (most connected - your core abstractions)
1. `AgentLoop` - 50 edges
2. `Store` - 48 edges
3. `Manager` - 41 edges
4. `newTestManager()` - 40 edges
5. `WeComChannel` - 33 edges
6. `newTestEngine()` - 29 edges
7. `turnState` - 29 edges
8. `openTestStore()` - 27 edges
9. `MatrixChannel` - 25 edges
10. `logMessage()` - 24 edges

## Surprising Connections (you probably didn't know these)
- `main()` --calls--> `NewPicoclawCommand()`  [INFERRED]
  examples/pico-echo-server/main.go → cmd/picoclaw/main.go
- `main()` --calls--> `shouldEnableLauncherFileLogging()`  [INFERRED]
  examples/pico-echo-server/main.go → web/backend/main.go
- `main()` --calls--> `dashboardTokenConfigHelpPath()`  [INFERRED]
  examples/pico-echo-server/main.go → web/backend/main.go
- `getGatewayStatus()` --calls--> `picoHome()`  [INFERRED]
  web/frontend/src/api/gateway.ts → cmd/picoclaw-launcher-tui/ui/gateway.go
- `Load()` --calls--> `DefaultConfig()`  [INFERRED]
  web/backend/launcherconfig/config.go → cmd/picoclaw-launcher-tui/config/config.go

## Communities

### Community 0 - "Tool Provider Interface"
Cohesion: 0.01
Nodes (86): fakeTranscriber, TestAgentCheckSilencePublishesInboundAndCleansUp(), waitForFileRemoval(), findDefinitionByName(), TestBuiltinHelpHandler_ReturnsFormattedMessage(), ClaudeProvider, formatHelpMessage(), helpCommand() (+78 more)

### Community 1 - "Channel Message Handler"
Cohesion: 0.01
Nodes (99): MessageBus, publish(), StreamDelegate, Streamer, buildChannelConfigResponse(), channelCatalogItem, channelConfigResponse, channelSecretPresence (+91 more)

### Community 2 - "Data Conversion Utils"
Cohesion: 0.01
Nodes (62): handleSave(), validate(), ansi256ToHex(), applyExtendedColor(), codeToColor(), parseAnsiSegments(), asString(), getFieldValueForValidation() (+54 more)

### Community 3 - "Access Control"
Cohesion: 0.01
Nodes (54): clientIPFromRemoteAddr(), IPAllowlist(), rejectByPolicy(), AnthropicUsage, loginRateLimiter, createAnthropicTestClient(), TestClaudeProvider_ChatRoundTrip(), createOpenAITestClient() (+46 more)

### Community 4 - "Provider API Config"
Cohesion: 0.01
Nodes (123): AgentBinding, AgentConfig, AgentDefaults, AgentModelConfig, AgentsConfig, applyConfigSecretsFromMap(), applyDefaults(), asMapField() (+115 more)

### Community 5 - "Audio Transcriber"
Cohesion: 0.01
Nodes (44): fakeLLMProvider, appendFile(), AppendFileTool, editFile(), EditFileTool, replaceEditContent(), assertLocalRateLimitSkipsToHealthyFallback(), makeCandidate() (+36 more)

### Community 6 - "Media Store Tools"
Cohesion: 0.02
Nodes (62): LoadImageTool, aesEcbPaddedSize(), buildCDNDownloadURL(), buildCDNUploadURL(), buildWeComVideoContent(), candidateWeComFilename(), canWeComSendImage(), canWeComSendVideo() (+54 more)

### Community 7 - "App Runtime Audio"
Cohesion: 0.02
Nodes (72): AudioModelTranscriber, botGoLogger, shouldDemoteBotGoInfo(), commandRegistrationDelay(), TelegramChannel, compressionResult, legacyContextManager, ElevenLabsTranscriber (+64 more)

### Community 8 - "LLM Provider Tools"
Cohesion: 0.02
Nodes (54): asyncFollowUpTool, collectEventStream(), findEvent(), scriptedToolProvider, stringError, TestAgentLoop_EmitsContextCompressEventOnRetry(), TestAgentLoop_EmitsFollowUpQueuedEvent(), TestAgentLoop_EmitsMinimalTurnEvents() (+46 more)

### Community 9 - "CLI SPI Provider"
Cohesion: 0.02
Nodes (43): createArgCaptureCLI(), createMockCLI(), createSlowMockCLI(), TestChat_ContextCancellation(), TestChat_EmptyWorkspaceDoesNotSetDir(), TestChat_InvalidResponseJSON(), TestChat_IsErrorResponse(), TestChat_NonZeroExitNoStderr() (+35 more)

### Community 10 - "Provider Core"
Cohesion: 0.02
Nodes (77): DecodeToolCallArguments(), ExpandHome(), extractCardImageKeys(), extractFileKey(), extractFileName(), extractImageKey(), extractImageKeysRecursive(), extractJSONStringField() (+69 more)

### Community 11 - "Agent Loop Provider"
Cohesion: 0.02
Nodes (43): artifactThenSendProvider, countingMockProvider, failFirstMockProvider, fakeChannel, fakeMediaChannel, handledMediaProvider, handledMediaTool, handledMediaWithSteeringProvider (+35 more)

### Community 12 - "Seahorse Memory"
Cohesion: 0.03
Nodes (60): newSeahorseContextManager(), providerToCompleteFn(), providerToSeahorseMessage(), seahorseContextManager, seahorseToProviderMessages(), AgentInstance, buildAllowReadPatterns(), compilePatterns() (+52 more)

### Community 13 - "Agent Command"
Cohesion: 0.03
Nodes (58): Agent, speechAccumulator, ApiClient, randomWechatUIN(), CronTool, JobExecutor, newTestCronTool(), newTestCronToolWithConfig() (+50 more)

### Community 14 - "Frontend Chat UI"
Cohesion: 0.02
Nodes (58): handleImageSelection(), handleScroll(), readFileAsDataUrl(), syncScrollState(), clearReconnectTimer(), connectChat(), disconnectChat(), disconnectChatInternal() (+50 more)

### Community 15 - "Auth Test Suite"
Cohesion: 0.02
Nodes (21): assertGatewayLogLevelApplied(), mustSetupSSHKey(), setupPicoEnabledEnv(), testCommandPatterns(), TestHandlePatchConfig_AppliesGatewayLogLevel(), TestHandlePatchConfig_SucceedsWhenPicoTokenInSecurityOnly(), TestHandleTestCommandPatterns_CaseInsensitiveWithGoFlag(), TestHandleTestCommandPatterns_EmptyPatterns() (+13 more)

### Community 16 - "USB Hardware Monitor"
Cohesion: 0.03
Nodes (47): Action, ContextCompressPayload, ContextCompressReason, DeviceEvent, ErrorPayload, Event, EventKind, EventMeta (+39 more)

### Community 17 - "CLI Install Commands"
Cohesion: 0.03
Nodes (55): buildOccupiedWorkspaceSkillsByDirectory(), buildSkillSupportItems(), buildWorkspaceSkillItemsByDirectory(), builtinSkillsDir(), commitStagedSkillInstall(), copyImportedSkillTree(), createStagedSkillInstall(), deleteSkill() (+47 more)

### Community 18 - "Launcher Auth"
Cohesion: 0.03
Nodes (44): AuthFlowOpts, launcherAuthHandlers, launcherAuthLoginBody, LauncherAuthRouteOpts, launcherAuthStatusResponse, LauncherAuthTokenHelp, buildCodexParams(), CodexProvider (+36 more)

### Community 19 - "Bedrock Provider Tests"
Cohesion: 0.02
Nodes (69): Action, ActionType, APIStatus, AudioChunk, BaseInfo, CacheControl, CDNMedia, ContentBlock (+61 more)

### Community 20 - "Config Management"
Cohesion: 0.04
Nodes (28): AuthCredential, authFilePath(), AuthStore, buildLikeQuery(), buildMessagesLikeQuery(), CleanupPolicy, contains(), CreateSummaryInput (+20 more)

### Community 21 - "Channel Handler"
Cohesion: 0.04
Nodes (39): escapeXML(), extractMarkdownMetadata(), nodeText(), SkillInfo, SkillMetadata, SkillsLoader, splitFrontmatter(), converter (+31 more)

### Community 22 - "Provider Tests"
Cohesion: 0.03
Nodes (63): AgentBinding, AgentConfig, AgentDefaults, AgentModelConfig, AgentsConfig, BindingMatch, BraveConfig, ChannelsConfig (+55 more)

### Community 23 - "Web API Backend"
Cohesion: 0.03
Nodes (44): mockAsyncRegistryTool, mockContextAwareTool, mockMediaStoreAwareTool, mockNilResultTool, mockPanicTool, mockRegistry, mockRegistryProvider, mockRegistryTool (+36 more)

### Community 24 - "Fyne UI Components"
Cohesion: 0.03
Nodes (33): chatWithCacheKey(), createAnthropicTestClient(), errAfterDataReadCloser, roundTripperFunc, TestProvider_ChatRoundTrip(), TestProviderChat_AcceptsNumericOptionTypes(), TestProviderChat_AzureAuthHeader(), TestProviderChat_AzureNativeWebSearchInjection() (+25 more)

### Community 25 - "Tool Execution"
Cohesion: 0.05
Nodes (33): applyWeComAuthResult(), authWeComCmd(), authWeComCmdWithScanner(), buildWeComQRCodePageURL(), buildWecomQRGenerateURL(), buildWecomQRQueryURL(), defaultWeComQRFlowOptions(), doWecomJSONGet() (+25 more)

### Community 26 - "Community 26"
Cohesion: 0.05
Nodes (57): App, applyGatewayStatusToStore(), attachToGatewayProcessLocked(), beginGatewayStoppingTransition(), cancelGatewayStoppingTransition(), canonicalGatewayLogLevel(), clearGatewayLogs(), clearGatewayStoppingTimeout() (+49 more)

### Community 27 - "Community 27"
Cohesion: 0.05
Nodes (29): DetectTranscriber(), fallbackTranscriberFromModelConfig(), supportsAudioTranscription(), supportsWhisperTranscription(), Transcriber, transcriberFromModelConfig(), TranscriptionResponse, whisperModelID() (+21 more)

### Community 28 - "Community 28"
Cohesion: 0.04
Nodes (46): mockChannel, mockChannelWithLength, mockDeletingMediaChannel, mockMediaChannel, mockMessageEditor, newTestManager(), TestInvokeTypingStop_CallsRegisteredStop(), TestInvokeTypingStop_Idempotent() (+38 more)

### Community 29 - "Community 29"
Cohesion: 0.04
Nodes (18): asyncTask, channelWorker, dispatchLoop(), finalizeHookStreamer, headerTransport, loadEnvFile(), Manager, newChannelWorker() (+10 more)

### Community 30 - "Community 30"
Cohesion: 0.04
Nodes (35): agentDefaultsV0, agentsConfigV0, baiduSearchConfigV0, braveConfigV0, channelsConfigV0, clawHubRegistryConfigV0, configV0, configV1 (+27 more)

### Community 31 - "Community 31"
Cohesion: 0.06
Nodes (33): appendUniquePath(), buildFs(), consumeNextLine(), containsUnescapedRegexMeta(), extractAllowedPathRoot(), fileSystem, formatDirEntries(), formatReadFileLinePrefix() (+25 more)

### Community 32 - "Community 32"
Cohesion: 0.05
Nodes (45): createTestZip(), newTestRegistry(), TestClawHubRegistryAuthToken(), TestClawHubRegistryDownloadAndInstall(), TestClawHubRegistryDownloadAndInstallRetries429(), TestClawHubRegistryGetSkillMeta(), TestClawHubRegistryGetSkillMetaUnsafeSlug(), TestClawHubRegistrySearch() (+37 more)

### Community 33 - "Community 33"
Cohesion: 0.06
Nodes (48): createTempFile(), newTestStoreWithCleanup(), openTestStore(), TestCleanExpiredCleansEmptyScopes(), TestCleanExpiredForgetOnlyKeepsFile(), TestCleanExpiredKeepsNonExpired(), TestCleanExpiredMixedAges(), TestCleanExpiredRemovesOldEntries() (+40 more)

### Community 34 - "Community 34"
Cohesion: 0.03
Nodes (0): 

### Community 35 - "Community 35"
Cohesion: 0.06
Nodes (42): base64URLDecode(), BuildAuthorizeURL(), buildOAuthRedirectURI(), callbackResult, decodeBase64(), defaultModelConfigForProvider(), DeviceCodeInfo, deviceCodeResponse (+34 more)

### Community 36 - "Community 36"
Cohesion: 0.06
Nodes (35): APIKeyIterator, APIKeyPool, BaiduSearchProvider, BraveSearchProvider, DuckDuckGoSearchProvider, GLMSearchProvider, isObviousPrivateHost(), isPrivateOrRestrictedIP() (+27 more)

### Community 37 - "Community 37"
Cohesion: 0.05
Nodes (20): AgentRegistry, ChannelFactory, ClawHubConfig, HiddenToolDoc, HiddenToolSnapshot, InstallResult, mediaStoreAware, NewRegistry() (+12 more)

### Community 38 - "Community 38"
Cohesion: 0.08
Nodes (25): ApprovalDecision, cloneLLMResponse(), cloneProviderMessages(), cloneProviderToolCalls(), cloneStringAnyMap(), cloneToolDefinitions(), cloneToolResult(), closeHookIfPossible() (+17 more)

### Community 39 - "Community 39"
Cohesion: 0.04
Nodes (10): blockingDirectProvider, capturingMockProvider, gracefulCaptureProvider, interruptibleTool, lateSteeringProvider, slowTool, TestAgentLoop_Steering_SkippedToolsHaveErrorResults(), toolCallProvider (+2 more)

### Community 40 - "Community 40"
Cohesion: 0.05
Nodes (17): serverHostAndPort(), singleHostCIDR(), TestWebFetch_RedirectToPrivateBlocked(), TestWebFetchTool_CloudflareChallenge_NoRetryOnOtherErrors(), TestWebFetchTool_CloudflareChallenge_RetryFailsToo(), TestWebFetchTool_CloudflareChallenge_RetryWithHonestUA(), TestWebFetchTool_PayloadTooLarge(), TestWebTool_WebFetch_HTMLExtraction() (+9 more)

### Community 41 - "Community 41"
Cohesion: 0.04
Nodes (0): 

### Community 42 - "Community 42"
Cohesion: 0.05
Nodes (8): eventCollector, newEventCollector(), panicMockProvider, simpleMockProviderAPI, slowMockProvider, TestSpawnSubTurn(), TestSpawnSubTurn_OrphanResultRouting(), TestSpawnSubTurn_PanicRecovery()

### Community 43 - "Community 43"
Cohesion: 0.04
Nodes (0): 

### Community 44 - "Community 44"
Cohesion: 0.08
Nodes (12): generateSecureToken(), getPicoToken(), Handler, inlineImageValue(), parseInlineImageMedia(), PicoChannel, picoConn, regenPicoToken() (+4 more)

### Community 45 - "Community 45"
Cohesion: 0.07
Nodes (26): assertHandleMessageQuotedUserReply(), multipartCall, multipartRecordingConstructor, newTestChannel(), newTestChannelWithConstructor(), stubCall, stubCaller, stubConstructor (+18 more)

### Community 46 - "Community 46"
Cohesion: 0.09
Nodes (33): newTestEngine(), newTestEngineForConcurrency(), TestAssemblerLazyInitRace(), TestAssemblerSummaryRoleNotUser(), TestBootstrapAnchorWithDuplicateContent(), TestBootstrapAnchorWithDuplicateContent_Simple(), TestBootstrapDeltaPreservesOrder(), TestBootstrapDuplicateContent() (+25 more)

### Community 47 - "Community 47"
Cohesion: 0.05
Nodes (0): 

### Community 48 - "Community 48"
Cohesion: 0.1
Nodes (14): cryptoRandInt(), isBotCommandEntityForThisBot(), isPostConnectError(), logParseFailed(), parseContent(), parseTelegramChatID(), quotedTelegramMediaRefs(), sendChunkParams (+6 more)

### Community 49 - "Community 49"
Cohesion: 0.06
Nodes (13): AsyncCallback, AsyncExecutor, BaseChannel, BaseChannelOption, BuildMediaScope(), Channel, MessageLengthProvider, Tool (+5 more)

### Community 50 - "Community 50"
Cohesion: 0.08
Nodes (12): deps, formatModelName(), listAvailableModels(), NewModelCommand(), NewVersionCommand(), printVersion(), setDefaultModel(), showCurrentModel() (+4 more)

### Community 51 - "Community 51"
Cohesion: 0.07
Nodes (9): TestMigrateFromJSON_Basic(), TestMigrateFromJSON_ColonInKey(), TestMigrateFromJSON_Idempotent(), TestMigrateFromJSON_InvalidJSON(), TestMigrateFromJSON_MultipleFiles(), TestMigrateFromJSON_RenamesFiles(), TestMigrateFromJSON_RetryAfterCrash(), TestMigrateFromJSON_WithToolCalls() (+1 more)

### Community 52 - "Community 52"
Cohesion: 0.07
Nodes (10): fixedScoreClassifier, TestRuleClassifier_AttachmentsHardGate(), TestRuleClassifier_CodeBlockAlone(), TestRuleClassifier_DeepConversation(), TestRuleClassifier_LongMessage(), TestRuleClassifier_MediumMessage(), TestRuleClassifier_ScoreDoesNotExceedOne(), TestRuleClassifier_ShortMessage() (+2 more)

### Community 53 - "Community 53"
Cohesion: 0.12
Nodes (27): apiRootFromAPIBase(), getJSON(), hasLocalAPIBase(), hasModelConfiguration(), hostPortFromAPIBase(), modelConfigurationStatus(), modelConfigurationSummary, modelProbeAPIBase() (+19 more)

### Community 54 - "Community 54"
Cohesion: 0.11
Nodes (4): AgentLoop, normalizeSteeringScope(), SteeringMode, steeringQueue

### Community 55 - "Community 55"
Cohesion: 0.11
Nodes (11): BotStatus, isAPIResponse(), oneBotAPIRequest, OneBotChannel, oneBotMessageSegment, oneBotRawEvent, oneBotSender, parseJSONInt64() (+3 more)

### Community 56 - "Community 56"
Cohesion: 0.12
Nodes (11): appendContent(), isHTTPURL(), qqAPI, qqAttachmentFilename(), qqAttachmentKind(), qqAttachmentNote(), QQChannel, qqFileType() (+3 more)

### Community 57 - "Community 57"
Cohesion: 0.06
Nodes (0): 

### Community 58 - "Community 58"
Cohesion: 0.15
Nodes (26): newCMTestAgentLoop(), noopContextManager, resetCMRegistry(), TestAgentLoop_UsesCustomContextManager(), testConfig(), TestIngestCalledDuringTurn(), TestLegacyAssemble_EmptyHistory(), TestLegacyAssemble_Passthrough() (+18 more)

### Community 59 - "Community 59"
Cohesion: 0.11
Nodes (17): AssembleInput, Assembler, AssembleResult, CompactionEngine, compileSessionPattern(), compileSessionPatterns(), CompleteFn, CompleteOptions (+9 more)

### Community 60 - "Community 60"
Cohesion: 0.1
Nodes (12): generateQRDataURI(), Handler, init(), NewWeixinChannel(), newWeixinCommand(), newWeixinFlowID(), printManualWeixinConfig(), runWeixinOnboard() (+4 more)

### Community 61 - "Community 61"
Cohesion: 0.11
Nodes (13): AgentLoop, NewProcessHook(), newProcessHookObserveKinds(), ProcessHook, processHookAfterLLMResponse, processHookAfterToolResponse, processHookBeforeLLMResponse, processHookBeforeToolResponse (+5 more)

### Community 62 - "Community 62"
Cohesion: 0.22
Nodes (23): assertRoles(), assistantWithTools(), containsSubstring(), findSubstring(), msg(), roles(), TestContextTruncationFlow(), TestSanitizeHistoryForProvider_AssistantToolCallAfterPlainAssistant() (+15 more)

### Community 63 - "Community 63"
Cohesion: 0.1
Nodes (23): antigravityContent, antigravityFuncDecl, antigravityFunctionCall, antigravityFunctionResponse, antigravityGenConfig, antigravityJSONResponse, AntigravityModelInfo, antigravityPart (+15 more)

### Community 64 - "Community 64"
Cohesion: 0.11
Nodes (7): buildTextMessage(), LINEChannel, lineEvent, lineMentionee, lineMessage, lineSource, replyTokenEntry

### Community 65 - "Community 65"
Cohesion: 0.13
Nodes (22): newTestStore(), TestAddFullMessage_ToolCallID(), TestAddFullMessage_WithToolCalls(), TestAddMessage_AutoCreatesSession(), TestAddMessage_BasicRoundtrip(), TestColonInKey(), TestCompact_NoOpWhenNoSkip(), TestCompact_RemovesSkippedMessages() (+14 more)

### Community 66 - "Community 66"
Cohesion: 0.1
Nodes (12): AgentLoop, AgentLoopFromContext(), agentLoopKeyType, AgentLoopSpawner, deliverSubTurnResult(), ephemeralSessionStore, ephemeralSessionStoreIface, newEphemeralSession() (+4 more)

### Community 67 - "Community 67"
Cohesion: 0.08
Nodes (3): mockOperation, TestMigrateInstanceGetCurrentHandler(), TestMigrateInstanceGetCurrentHandlerWithSource()

### Community 68 - "Community 68"
Cohesion: 0.11
Nodes (9): callerFromYaml(), FlexibleStringSlice, NewSecureString(), resolveKey(), SecureModelList, SecureString, SecureStrings, SimpleSecureStrings() (+1 more)

### Community 69 - "Community 69"
Cohesion: 0.12
Nodes (12): captureStdout(), initTest(), TestListAvailableModels_Empty(), TestListAvailableModels_MarkerLogic(), TestListAvailableModels_WithModels(), TestModelCommandExecution_Set(), TestModelCommandExecution_Show(), TestSetDefaultModel_InvalidModel() (+4 more)

### Community 70 - "Community 70"
Cohesion: 0.15
Nodes (25): newTestCompactionEngine(), newTestCompactionEngineWithStore(), TestCompactAsyncDedup(), TestCompactAsyncReturnsBeforeCondensed(), TestCompactCondensed(), TestCompactCondensedDoesNotOrphanSummaryWhenCandidatesRemovedConcurrently(), TestCompactCondensedUsesOrdinalAwareSelection(), TestCompactCondensedUsesSelectOldestChunk() (+17 more)

### Community 71 - "Community 71"
Cohesion: 0.08
Nodes (1): MockMCPManager

### Community 72 - "Community 72"
Cohesion: 0.08
Nodes (3): containsString(), searchString(), TestLoadConfig_UnsupportedVersion()

### Community 73 - "Community 73"
Cohesion: 0.08
Nodes (1): seahorseTestProvider

### Community 74 - "Community 74"
Cohesion: 0.11
Nodes (7): cachedSystemVersion, fallbackSystemVersionInfo(), Handler, isLikelyVersionValue(), parsePicoclawVersionOutput(), systemVersionCache, systemVersionResponse

### Community 75 - "Community 75"
Cohesion: 0.15
Nodes (13): buildAggressiveLeafSummaryPrompt(), buildCondensedSummaryPrompt(), buildLeafSummaryPrompt(), CompactInput, CompactionEngine, CompactResult, formatMessagesForSummary(), formatSummariesForCondensation() (+5 more)

### Community 76 - "Community 76"
Cohesion: 0.08
Nodes (6): SpawnSubTurnFunc, SubagentManager, SubagentTask, SubagentTool, SubTurnConfig, SubTurnSpawner

### Community 77 - "Community 77"
Cohesion: 0.14
Nodes (16): assertAudioWAVUploadType(), fakeQQAPI, fakeTransportCall, mustJSON(), TestHandleC2CMessage_AttachmentOnlyPublishesMedia(), TestHandleGroupATMessage_AttachmentOnlyPublishesMedia(), TestSendMedia_AudioAt60SecondsUsesVoiceUpload(), TestSendMedia_AudioOver60SecondsFallsBackToFileUpload() (+8 more)

### Community 78 - "Community 78"
Cohesion: 0.22
Nodes (18): msgAssistant(), msgAssistantTC(), msgTool(), msgUser(), TestEstimateMessageTokens(), TestEstimateMessageTokens_MediaItems(), TestEstimateMessageTokens_MultibyteContent(), TestEstimateMessageTokens_ReasoningContent() (+10 more)

### Community 79 - "Community 79"
Cohesion: 0.12
Nodes (9): bm25CachedEngine, BM25SearchTool, buildBM25Engine(), formatDiscoveryResponse(), RegexSearchTool, searchDoc, snapshotToSearchDocs(), ToolRegistry (+1 more)

### Community 80 - "Community 80"
Cohesion: 0.17
Nodes (9): annotationsAllowUser(), compactStrings(), extractContentText(), MCPManager, MCPTool, normalizedMIMEType(), sanitizeIdentifierComponent(), summarizeEmbeddedResource() (+1 more)

### Community 81 - "Community 81"
Cohesion: 0.16
Nodes (15): autoStartRequest, autoStartResponse, buildDarwinPlist(), buildLinuxExecLine(), fileExists(), Handler, linuxAutoStartPath(), macLaunchAgentPath() (+7 more)

### Community 82 - "Community 82"
Cohesion: 0.12
Nodes (8): mustGatewayTestPort(), newTestPicoChannel(), PicoChannel, TestBroadcastToSession_TargetsOnlyRequestedSession(), TestCreateAndAddConnection_RespectsMaxConnectionsConcurrently(), TestHandleWebSocketProxyLoadsCachedPicoTokenWhenMissing(), TestHandleWebSocketProxyReloadsGatewayTargetFromConfig(), TestRemoveConnection_CleansBothIndexes()

### Community 83 - "Community 83"
Cohesion: 0.12
Nodes (16): AgentContextDefinition, AgentDefinitionSource, AgentFrontmatter, AgentPromptDefinition, ContextBuilder, Definition, fileExists(), loadAgentDefinition() (+8 more)

### Community 84 - "Community 84"
Cohesion: 0.15
Nodes (12): AgentLoop, BuiltinHookFactory, configureHookManagerFromConfig(), enabledBuiltinHookNames(), enabledProcessHookNames(), hookRuntime, hookTimeoutFromMS(), lookupBuiltinHook() (+4 more)

### Community 85 - "Community 85"
Cohesion: 0.2
Nodes (10): clientVisiblePort(), forwardedHostFirst(), forwardedPortFirst(), forwardedRFC7239Host(), gatewayProbeHost(), Handler, joinClientVisibleHostPort(), requestHostName() (+2 more)

### Community 86 - "Community 86"
Cohesion: 0.12
Nodes (5): boolPtr(), strPtr(), TestChannelEnabled(), TestGetDefaultModel(), TestOpenClawAgentModel()

### Community 87 - "Community 87"
Cohesion: 0.18
Nodes (4): App, cacheKey(), hintBar(), New()

### Community 88 - "Community 88"
Cohesion: 0.12
Nodes (6): Check, extractBearerToken(), HandlerMux, Server, StatusResponse, statusString()

### Community 89 - "Community 89"
Cohesion: 0.15
Nodes (14): bm25Config, bm25Dedupe(), bm25DocEntry, BM25Engine, BM25Engine[T], bm25Index, bm25MinHeapify(), BM25Option (+6 more)

### Community 90 - "Community 90"
Cohesion: 0.11
Nodes (0): 

### Community 91 - "Community 91"
Cohesion: 0.15
Nodes (7): createSkillDir(), TestListSkillsDirWithoutSkillMD(), TestListSkillsGlobalOverridesBuiltin(), TestListSkillsInvalidSkillSkipped(), TestListSkillsMetadataNameDedup(), TestListSkillsMultipleDistinctSkills(), TestListSkillsWorkspaceOverridesGlobal()

### Community 92 - "Community 92"
Cohesion: 0.18
Nodes (8): mockSearchableTool, setupPopulatedRegistry(), TestBM25CacheInvalidation(), TestBM25SearchTool_Execute(), TestRegexSearchTool_Execute(), TestRegexSearchTool_PatternTooLong(), TestSearchBM25_ZeroMaxResults(), TestSearchRegex_ZeroMaxResults()

### Community 93 - "Community 93"
Cohesion: 0.2
Nodes (2): Logger, maskSecrets()

### Community 94 - "Community 94"
Cohesion: 0.12
Nodes (0): 

### Community 95 - "Community 95"
Cohesion: 0.21
Nodes (4): matchesAccountID(), ResolvedRoute, RouteInput, RouteResolver

### Community 96 - "Community 96"
Cohesion: 0.22
Nodes (12): setupAssemblerStore(), TestAssemblerAssembleEmpty(), TestAssemblerAssembleMessagesOnly(), TestAssemblerAssembleWithSummary(), TestAssemblerBudgetEvictsOldest(), TestAssemblerBudgetFitsAll(), TestAssemblerDepthAwarePrompt(), TestAssemblerLeafSummaryNoParents() (+4 more)

### Community 97 - "Community 97"
Cohesion: 0.12
Nodes (1): MockLLMProvider

### Community 98 - "Community 98"
Cohesion: 0.23
Nodes (14): setupWorkspace(), TestBuildMessages_CurrentSenderDynamicContext(), TestBuildMessages_IncludesMediaOnlyCurrentMessage(), TestBuiltinSkillFileContentChange(), TestCacheStability(), TestConcurrentBuildSystemPromptWithCache(), TestEmptyWorkspaceBaselineDetectsNewFiles(), TestExplicitInvalidateCache() (+6 more)

### Community 99 - "Community 99"
Cohesion: 0.23
Nodes (13): TestPidFilePath(), TestReadPidFileUnlockedInvalidJSON(), TestReadPidFileUnlockedInvalidPID(), TestReadPidFileWithCheck(), TestReadPidFileWithCheckNonexistent(), TestReadPidFileWithCheckStalePID(), TestRemovePidFile(), TestRemovePidFileDifferentPID() (+5 more)

### Community 100 - "Community 100"
Cohesion: 0.18
Nodes (6): newTestTracker(), TestCooldown_BillingEscalation(), TestCooldown_BillingTakesPrecedence(), TestCooldown_CooldownRemaining(), TestCooldown_FailureWindowReset(), TestCooldown_StandardEscalation()

### Community 101 - "Community 101"
Cohesion: 0.22
Nodes (6): clawhubModerationInfo, ClawHubRegistry, clawhubSearchResponse, clawhubSearchResult, clawhubSkillResponse, clawhubVersionInfo

### Community 102 - "Community 102"
Cohesion: 0.29
Nodes (13): BenchmarkAssemble_BudgetEviction(), BenchmarkAssemble_MessagesOnly(), BenchmarkAssemble_WithSummaries(), BenchmarkBootstrap_100Messages(), BenchmarkBootstrap_500Messages(), BenchmarkBootstrap_Empty(), BenchmarkIngest_BatchMessages(), BenchmarkIngest_SingleMessage() (+5 more)

### Community 103 - "Community 103"
Cohesion: 0.14
Nodes (0): 

### Community 104 - "Community 104"
Cohesion: 0.22
Nodes (3): claudeCliJSONResponse, ClaudeCliProvider, claudeCliUsageInfo

### Community 105 - "Community 105"
Cohesion: 0.19
Nodes (4): calculateBillingCooldown(), calculateStandardCooldown(), cooldownEntry, CooldownTracker

### Community 106 - "Community 106"
Cohesion: 0.21
Nodes (7): FallbackAttempt, FallbackCandidate, FallbackChain, FallbackExhaustedError, FallbackResult, ResolveCandidates(), ResolveCandidatesWithLookup()

### Community 107 - "Community 107"
Cohesion: 0.18
Nodes (4): benchmarkBM25Corpus(), BenchmarkBM25Search_RebuildEachTime(), BenchmarkBM25Search_ReusedIndex(), testDoc

### Community 108 - "Community 108"
Cohesion: 0.27
Nodes (10): blockedToolProvider, handleProcessHookRequest(), processHookHelperCommand(), processHookHelperEnv(), runProcessHookHelper(), TestAgentLoop_MountProcessHook_ApprovalDeny(), TestAgentLoop_MountProcessHook_LLMAndObserver(), TestAgentLoop_MountProcessHook_ToolRewrite() (+2 more)

### Community 109 - "Community 109"
Cohesion: 0.26
Nodes (9): cleanupWorkspace(), TestLoadAgentDefinitionFallsBackToLegacyAgentsMarkdown(), TestLoadAgentDefinitionInvalidFrontmatterFallsBackToEmptyStructuredFields(), TestLoadAgentDefinitionLoadsWorkspaceUserMarkdown(), TestLoadAgentDefinitionParsesFrontmatterAndSoul(), TestLoadBootstrapFilesIncludesWorkspaceUserMarkdown(), TestLoadBootstrapFilesUsesAgentBodyNotFrontmatter(), TestStructuredAgentIgnoresIdentityChanges() (+1 more)

### Community 110 - "Community 110"
Cohesion: 0.35
Nodes (6): buildTrigrams(), cacheEntry, copyResults(), jaccardSimilarity(), normalizeQuery(), SearchCache

### Community 111 - "Community 111"
Cohesion: 0.17
Nodes (1): ToolResult

### Community 112 - "Community 112"
Cohesion: 0.17
Nodes (0): 

### Community 113 - "Community 113"
Cohesion: 0.23
Nodes (3): newRateLimiter(), RateLimiter, RateLimiterRegistry

### Community 114 - "Community 114"
Cohesion: 0.2
Nodes (3): TestParseResponse_NoThinkingBlock(), TestParseResponse_ThinkingBlock(), unmarshalBlocks()

### Community 115 - "Community 115"
Cohesion: 0.32
Nodes (11): testConfig(), TestResolveRoute_AccountBinding(), TestResolveRoute_ChannelWildcard(), TestResolveRoute_DefaultAgent_NoBindings(), TestResolveRoute_DefaultAgentSelection(), TestResolveRoute_GuildBinding(), TestResolveRoute_InvalidAgentFallsToDefault(), TestResolveRoute_NoDefaultUsesFirst() (+3 more)

### Community 116 - "Community 116"
Cohesion: 0.18
Nodes (0): 

### Community 117 - "Community 117"
Cohesion: 0.22
Nodes (5): CodexCliProvider, codexEvent, codexEventErr, codexEventItem, codexUsage

### Community 118 - "Community 118"
Cohesion: 0.25
Nodes (10): BuildAgentMainSessionKey(), BuildAgentPeerSessionKey(), DMScope, IsSubagentSessionKey(), normalizeChannel(), ParseAgentSessionKey(), ParsedSessionKey, resolveLinkedPeerID() (+2 more)

### Community 119 - "Community 119"
Cohesion: 0.2
Nodes (0): 

### Community 120 - "Community 120"
Cohesion: 0.24
Nodes (3): collectSensitive(), Config, SensitiveDataCache

### Community 121 - "Community 121"
Cohesion: 0.2
Nodes (0): 

### Community 122 - "Community 122"
Cohesion: 0.24
Nodes (3): EventBus, eventSubscriber, EventSubscription

### Community 123 - "Community 123"
Cohesion: 0.2
Nodes (0): 

### Community 124 - "Community 124"
Cohesion: 0.36
Nodes (7): BuildMultipartContent(), ParseDataAudioURL(), parseResponse(), ParseResponseBody(), ParseResponseFromStruct(), ResolveToolCall(), TranslateMessages()

### Community 125 - "Community 125"
Cohesion: 0.22
Nodes (0): 

### Community 126 - "Community 126"
Cohesion: 0.33
Nodes (8): codeBlockMatch, escapeHTML(), extractCodeBlocks(), extractInlineCodes(), extractLinks(), inlineCodeMatch, linkMatch, markdownToTelegramHTML()

### Community 127 - "Community 127"
Cohesion: 0.46
Nodes (7): openTestDB(), TestFTS5SQLConstants(), TestMigrationConversationUnique(), TestMigrationSummaryFTSInsert(), TestMigrationSummaryParentsPK(), TestRunMigrations(), TestRunMigrationsIdempotent()

### Community 128 - "Community 128"
Cohesion: 0.25
Nodes (4): i2cSmbusArgs, i2cSmbusData, I2CTool, smbusProbe()

### Community 129 - "Community 129"
Cohesion: 0.25
Nodes (0): 

### Community 130 - "Community 130"
Cohesion: 0.25
Nodes (0): 

### Community 131 - "Community 131"
Cohesion: 0.48
Nodes (1): Handler

### Community 132 - "Community 132"
Cohesion: 0.52
Nodes (6): allowsAdditional(), checkArrayItems(), checkEnum(), checkRequired(), checkType(), validateToolArgs()

### Community 133 - "Community 133"
Cohesion: 0.43
Nodes (5): EstimateMessageTokens(), EstimateToolDefsTokens(), findSafeBoundary(), isOverContextBudget(), parseTurnBoundaries()

### Community 134 - "Community 134"
Cohesion: 0.48
Nodes (6): countCodeBlocks(), countRecentToolCalls(), estimateTokens(), ExtractFeatures(), Features, hasAttachments()

### Community 135 - "Community 135"
Cohesion: 0.52
Nodes (6): findLastNewlineInRange(), findLastSpaceInRange(), findLastUnclosedCodeBlockInRange(), findNewlineFrom(), findNextClosingCodeBlockInRange(), SplitMessage()

### Community 136 - "Community 136"
Cohesion: 0.6
Nodes (5): qqAudioDuration(), qqAudioDurationFormat(), qqOggDuration(), qqParseOggCodec(), qqWAVDuration()

### Community 137 - "Community 137"
Cohesion: 0.4
Nodes (0): 

### Community 138 - "Community 138"
Cohesion: 0.4
Nodes (0): 

### Community 139 - "Community 139"
Cohesion: 0.6
Nodes (4): ModelKey(), ModelRef, NormalizeProvider(), ParseModelRef()

### Community 140 - "Community 140"
Cohesion: 0.4
Nodes (0): 

### Community 141 - "Community 141"
Cohesion: 0.7
Nodes (1): DiscordChannel

### Community 142 - "Community 142"
Cohesion: 0.7
Nodes (4): entityPattern, escapeMarkdownV2(), markdownToTelegramMarkdownV2(), processText()

### Community 143 - "Community 143"
Cohesion: 0.5
Nodes (0): 

### Community 144 - "Community 144"
Cohesion: 0.5
Nodes (1): I2CTool

### Community 145 - "Community 145"
Cohesion: 0.5
Nodes (0): 

### Community 146 - "Community 146"
Cohesion: 0.67
Nodes (2): buildOggPage(), TestDecodeOggOpus_ValidParsing()

### Community 147 - "Community 147"
Cohesion: 0.5
Nodes (2): Classifier, RuleClassifier

### Community 148 - "Community 148"
Cohesion: 0.67
Nodes (3): DetectVoiceCapabilities(), VoiceCapabilities, VoiceCapabilityProvider

### Community 149 - "Community 149"
Cohesion: 0.5
Nodes (0): 

### Community 150 - "Community 150"
Cohesion: 0.5
Nodes (0): 

### Community 151 - "Community 151"
Cohesion: 1.0
Nodes (1): App

### Community 152 - "Community 152"
Cohesion: 1.0
Nodes (1): App

### Community 153 - "Community 153"
Cohesion: 1.0
Nodes (1): App

### Community 154 - "Community 154"
Cohesion: 0.67
Nodes (0): 

### Community 155 - "Community 155"
Cohesion: 0.67
Nodes (0): 

### Community 156 - "Community 156"
Cohesion: 0.67
Nodes (0): 

### Community 157 - "Community 157"
Cohesion: 0.67
Nodes (1): Handler

### Community 158 - "Community 158"
Cohesion: 0.67
Nodes (0): 

### Community 159 - "Community 159"
Cohesion: 0.67
Nodes (0): 

### Community 160 - "Community 160"
Cohesion: 0.67
Nodes (1): SPITool

### Community 161 - "Community 161"
Cohesion: 0.67
Nodes (0): 

### Community 162 - "Community 162"
Cohesion: 1.0
Nodes (2): CopyFile(), WriteFileAtomic()

### Community 163 - "Community 163"
Cohesion: 1.0
Nodes (2): InitPanic(), RecoverPanicNoExit()

### Community 164 - "Community 164"
Cohesion: 0.67
Nodes (0): 

### Community 165 - "Community 165"
Cohesion: 0.67
Nodes (1): PKCECodes

### Community 166 - "Community 166"
Cohesion: 0.67
Nodes (0): 

### Community 167 - "Community 167"
Cohesion: 0.67
Nodes (0): 

### Community 168 - "Community 168"
Cohesion: 0.67
Nodes (0): 

### Community 169 - "Community 169"
Cohesion: 0.67
Nodes (0): 

### Community 170 - "Community 170"
Cohesion: 1.0
Nodes (2): matchesMagic(), TestDownloadAndExtractRelease_RealPlatforms()

### Community 171 - "Community 171"
Cohesion: 0.67
Nodes (1): ThinkingLevel

### Community 172 - "Community 172"
Cohesion: 1.0
Nodes (2): mergeShorties(), SplitSentences()

### Community 173 - "Community 173"
Cohesion: 0.67
Nodes (0): 

### Community 174 - "Community 174"
Cohesion: 1.0
Nodes (1): App

### Community 175 - "Community 175"
Cohesion: 1.0
Nodes (0): 

### Community 176 - "Community 176"
Cohesion: 1.0
Nodes (0): 

### Community 177 - "Community 177"
Cohesion: 1.0
Nodes (0): 

### Community 178 - "Community 178"
Cohesion: 1.0
Nodes (0): 

### Community 179 - "Community 179"
Cohesion: 1.0
Nodes (0): 

### Community 180 - "Community 180"
Cohesion: 1.0
Nodes (0): 

### Community 181 - "Community 181"
Cohesion: 1.0
Nodes (0): 

### Community 182 - "Community 182"
Cohesion: 1.0
Nodes (0): 

### Community 183 - "Community 183"
Cohesion: 1.0
Nodes (0): 

### Community 184 - "Community 184"
Cohesion: 1.0
Nodes (0): 

### Community 185 - "Community 185"
Cohesion: 1.0
Nodes (0): 

### Community 186 - "Community 186"
Cohesion: 1.0
Nodes (0): 

### Community 187 - "Community 187"
Cohesion: 1.0
Nodes (0): 

### Community 188 - "Community 188"
Cohesion: 1.0
Nodes (0): 

### Community 189 - "Community 189"
Cohesion: 1.0
Nodes (0): 

### Community 190 - "Community 190"
Cohesion: 1.0
Nodes (0): 

### Community 191 - "Community 191"
Cohesion: 1.0
Nodes (0): 

### Community 192 - "Community 192"
Cohesion: 1.0
Nodes (0): 

### Community 193 - "Community 193"
Cohesion: 1.0
Nodes (0): 

### Community 194 - "Community 194"
Cohesion: 1.0
Nodes (0): 

### Community 195 - "Community 195"
Cohesion: 1.0
Nodes (0): 

### Community 196 - "Community 196"
Cohesion: 1.0
Nodes (0): 

### Community 197 - "Community 197"
Cohesion: 1.0
Nodes (0): 

### Community 198 - "Community 198"
Cohesion: 1.0
Nodes (0): 

### Community 199 - "Community 199"
Cohesion: 1.0
Nodes (0): 

### Community 200 - "Community 200"
Cohesion: 1.0
Nodes (0): 

### Community 201 - "Community 201"
Cohesion: 1.0
Nodes (0): 

### Community 202 - "Community 202"
Cohesion: 1.0
Nodes (0): 

### Community 203 - "Community 203"
Cohesion: 1.0
Nodes (0): 

### Community 204 - "Community 204"
Cohesion: 1.0
Nodes (0): 

### Community 205 - "Community 205"
Cohesion: 1.0
Nodes (0): 

### Community 206 - "Community 206"
Cohesion: 1.0
Nodes (0): 

### Community 207 - "Community 207"
Cohesion: 1.0
Nodes (0): 

### Community 208 - "Community 208"
Cohesion: 1.0
Nodes (0): 

### Community 209 - "Community 209"
Cohesion: 1.0
Nodes (0): 

### Community 210 - "Community 210"
Cohesion: 1.0
Nodes (0): 

### Community 211 - "Community 211"
Cohesion: 1.0
Nodes (0): 

### Community 212 - "Community 212"
Cohesion: 1.0
Nodes (0): 

### Community 213 - "Community 213"
Cohesion: 1.0
Nodes (0): 

### Community 214 - "Community 214"
Cohesion: 1.0
Nodes (0): 

### Community 215 - "Community 215"
Cohesion: 1.0
Nodes (0): 

### Community 216 - "Community 216"
Cohesion: 1.0
Nodes (0): 

### Community 217 - "Community 217"
Cohesion: 1.0
Nodes (0): 

### Community 218 - "Community 218"
Cohesion: 1.0
Nodes (1): FeishuChannel

### Community 219 - "Community 219"
Cohesion: 1.0
Nodes (0): 

### Community 220 - "Community 220"
Cohesion: 1.0
Nodes (0): 

### Community 221 - "Community 221"
Cohesion: 1.0
Nodes (0): 

### Community 222 - "Community 222"
Cohesion: 1.0
Nodes (1): SessionStore

### Community 223 - "Community 223"
Cohesion: 1.0
Nodes (0): 

### Community 224 - "Community 224"
Cohesion: 1.0
Nodes (0): 

### Community 225 - "Community 225"
Cohesion: 1.0
Nodes (0): 

### Community 226 - "Community 226"
Cohesion: 1.0
Nodes (0): 

### Community 227 - "Community 227"
Cohesion: 1.0
Nodes (0): 

### Community 228 - "Community 228"
Cohesion: 1.0
Nodes (0): 

### Community 229 - "Community 229"
Cohesion: 1.0
Nodes (0): 

### Community 230 - "Community 230"
Cohesion: 1.0
Nodes (0): 

### Community 231 - "Community 231"
Cohesion: 1.0
Nodes (0): 

## Knowledge Gaps
- **544 isolated node(s):** `modelsAPIResponse`, `modelEntry`, `App`, `gatewayStatus`, `App` (+539 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **Thin community `Community 174`** (2 nodes): `App`, `.newHomePage()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 175`** (2 nodes): `logout.go`, `newLogoutCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 176`** (2 nodes): `login_test.go`, `TestNewLoginSubCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 177`** (2 nodes): `logout_test.go`, `TestNewLogoutSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 178`** (2 nodes): `status_test.go`, `TestNewStatusSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 179`** (2 nodes): `login.go`, `newLoginCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 180`** (2 nodes): `remove_test.go`, `TestNewRemoveSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 181`** (2 nodes): `installbuiltin_test.go`, `TestNewInstallbuiltinSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 182`** (2 nodes): `search.go`, `newSearchCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 183`** (2 nodes): `listbuiltin_test.go`, `TestNewListbuiltinSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 184`** (2 nodes): `search_test.go`, `TestNewSearchSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 185`** (2 nodes): `installbuiltin.go`, `newInstallBuiltinCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 186`** (2 nodes): `listbuiltin.go`, `newListBuiltinCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 187`** (2 nodes): `show_test.go`, `TestNewShowSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 188`** (2 nodes): `list_test.go`, `TestNewListSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 189`** (2 nodes): `enable_test.go`, `TestEnableSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 190`** (2 nodes): `disable_test.go`, `TestDisableSubcommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 191`** (2 nodes): `disable.go`, `newDisableCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 192`** (2 nodes): `enable.go`, `newEnableCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 193`** (2 nodes): `CaptureButton.tsx`, `CaptureButton()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 194`** (2 nodes): `CameraView.tsx`, `CameraView()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 195`** (2 nodes): `VoiceButton.tsx`, `VoiceButton()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 196`** (2 nodes): `useVoice.ts`, `useVoice()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 197`** (2 nodes): `useCamera.ts`, `useCamera()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 198`** (2 nodes): `usePermissions.ts`, `usePermissions()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 199`** (2 nodes): `tray_offers_copy_stub.go`, `trayOffersDashboardTokenCopy()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 200`** (2 nodes): `tray_offers_copy.go`, `trayOffersDashboardTokenCopy()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 201`** (2 nodes): `session_process_windows.go`, `killProcessGroup()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 202`** (2 nodes): `sysproc_unix.go`, `setSysProcAttrForPty()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 203`** (2 nodes): `sysproc_windows.go`, `setSysProcAttrForPty()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 204`** (2 nodes): `session_process_unix.go`, `killProcessGroup()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 205`** (2 nodes): `panic_win.go`, `initPanicFile()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 206`** (2 nodes): `panic_unix.go`, `initPanicFile()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 207`** (2 nodes): `envkeys.go`, `GetHome()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 208`** (2 nodes): `defaults.go`, `DefaultConfig()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 209`** (2 nodes): `config_struct_test.go`, `TestLoadSecurityValue()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 210`** (2 nodes): `ogg.go`, `DecodeOggOpus()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 211`** (2 nodes): `sentence_test.go`, `TestSplitSentences()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 212`** (2 nodes): `pidfile_windows.go`, `isProcessRunning()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 213`** (2 nodes): `pidfile_unix.go`, `isProcessRunning()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 214`** (2 nodes): `request_test.go`, `TestHasCommandPrefix()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 215`** (2 nodes): `cmd_use.go`, `useCommand()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 216`** (2 nodes): `builtin.go`, `BuiltinDefinitions()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 217`** (2 nodes): `marker.go`, `SplitByMarker()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 218`** (2 nodes): `FeishuChannel`, `.VoiceCapabilities()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 219`** (2 nodes): `reqid_store_test.go`, `TestReqIDStorePersistsRoutes()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 220`** (2 nodes): `parser_markdown_to_html_test.go`, `Test_markdownToTelegramHTML()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 221`** (2 nodes): `tempdir.go`, `TempDir()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 222`** (2 nodes): `session_store.go`, `SessionStore`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 223`** (1 nodes): `home.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 224`** (1 nodes): `prettier.config.js`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 225`** (1 nodes): `vite.config.ts`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 226`** (1 nodes): `banner.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 227`** (1 nodes): `env.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 228`** (1 nodes): `short_constants.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 229`** (1 nodes): `i2c_other.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 230`** (1 nodes): `spi_other.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 231`** (1 nodes): `example_security_usage.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Are the 39 inferred relationships involving `newTestManager()` (e.g. with `TestStartAll_AllChannelsFail_ReturnsJoinedError()` and `TestStartAll_PartialFailure_StartsSuccessfulWorkers()`) actually correct?**
  _`newTestManager()` has 39 INFERRED edges - model-reasoned connections that need verification._
- **What connects `modelsAPIResponse`, `modelEntry`, `App` to the rest of the system?**
  _544 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Tool Provider Interface` be split into smaller, more focused modules?**
  _Cohesion score 0.01 - nodes in this community are weakly interconnected._
- **Should `Channel Message Handler` be split into smaller, more focused modules?**
  _Cohesion score 0.01 - nodes in this community are weakly interconnected._
- **Should `Data Conversion Utils` be split into smaller, more focused modules?**
  _Cohesion score 0.01 - nodes in this community are weakly interconnected._
- **Should `Access Control` be split into smaller, more focused modules?**
  _Cohesion score 0.01 - nodes in this community are weakly interconnected._
- **Should `Provider API Config` be split into smaller, more focused modules?**
  _Cohesion score 0.01 - nodes in this community are weakly interconnected._