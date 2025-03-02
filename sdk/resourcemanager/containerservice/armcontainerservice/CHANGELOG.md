# Release History

## 2.4.0-beta.1 (2023-01-27)
### Features Added

- New value `OSSKUMariner` added to type alias `OSSKU`
- New value `PublicNetworkAccessSecuredByPerimeter` added to type alias `PublicNetworkAccess`
- New value `SnapshotTypeManagedCluster` added to type alias `SnapshotType`
- New value `WorkloadRuntimeKataMshvVMIsolation` added to type alias `WorkloadRuntime`
- New type alias `BackendPoolType` with values `BackendPoolTypeNodeIP`, `BackendPoolTypeNodeIPConfiguration`
- New type alias `ControlledValues` with values `ControlledValuesRequestsAndLimits`, `ControlledValuesRequestsOnly`
- New type alias `EbpfDataplane` with values `EbpfDataplaneCilium`
- New type alias `FleetMemberProvisioningState` with values `FleetMemberProvisioningStateCanceled`, `FleetMemberProvisioningStateFailed`, `FleetMemberProvisioningStateJoining`, `FleetMemberProvisioningStateLeaving`, `FleetMemberProvisioningStateSucceeded`, `FleetMemberProvisioningStateUpdating`
- New type alias `FleetProvisioningState` with values `FleetProvisioningStateCanceled`, `FleetProvisioningStateCreating`, `FleetProvisioningStateDeleting`, `FleetProvisioningStateFailed`, `FleetProvisioningStateSucceeded`, `FleetProvisioningStateUpdating`
- New type alias `IpvsScheduler` with values `IpvsSchedulerLeastConnection`, `IpvsSchedulerRoundRobin`
- New type alias `Level` with values `LevelEnforcement`, `LevelOff`, `LevelWarning`
- New type alias `Mode` with values `ModeIPTABLES`, `ModeIPVS`
- New type alias `NetworkPluginMode` with values `NetworkPluginModeOverlay`
- New type alias `NodeOSUpgradeChannel` with values `NodeOSUpgradeChannelNodeImage`, `NodeOSUpgradeChannelNone`, `NodeOSUpgradeChannelSecurityPatch`, `NodeOSUpgradeChannelUnmanaged`
- New type alias `Protocol` with values `ProtocolTCP`, `ProtocolUDP`
- New type alias `RestrictionLevel` with values `RestrictionLevelReadOnly`, `RestrictionLevelUnrestricted`
- New type alias `TrustedAccessRoleBindingProvisioningState` with values `TrustedAccessRoleBindingProvisioningStateCanceled`, `TrustedAccessRoleBindingProvisioningStateDeleting`, `TrustedAccessRoleBindingProvisioningStateFailed`, `TrustedAccessRoleBindingProvisioningStateSucceeded`, `TrustedAccessRoleBindingProvisioningStateUpdating`
- New type alias `Type` with values `TypeFirst`, `TypeFourth`, `TypeLast`, `TypeSecond`, `TypeThird`
- New type alias `UpdateMode` with values `UpdateModeAuto`, `UpdateModeInitial`, `UpdateModeOff`, `UpdateModeRecreate`
- New function `*AgentPoolsClient.BeginAbortLatestOperation(context.Context, string, string, string, *AgentPoolsClientBeginAbortLatestOperationOptions) (*runtime.Poller[AgentPoolsClientAbortLatestOperationResponse], error)`
- New function `NewFleetMembersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FleetMembersClient, error)`
- New function `*FleetMembersClient.BeginCreateOrUpdate(context.Context, string, string, string, FleetMember, *FleetMembersClientBeginCreateOrUpdateOptions) (*runtime.Poller[FleetMembersClientCreateOrUpdateResponse], error)`
- New function `*FleetMembersClient.BeginDelete(context.Context, string, string, string, *FleetMembersClientBeginDeleteOptions) (*runtime.Poller[FleetMembersClientDeleteResponse], error)`
- New function `*FleetMembersClient.Get(context.Context, string, string, string, *FleetMembersClientGetOptions) (FleetMembersClientGetResponse, error)`
- New function `*FleetMembersClient.NewListByFleetPager(string, string, *FleetMembersClientListByFleetOptions) *runtime.Pager[FleetMembersClientListByFleetResponse]`
- New function `NewFleetsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FleetsClient, error)`
- New function `*FleetsClient.BeginCreateOrUpdate(context.Context, string, string, Fleet, *FleetsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FleetsClientCreateOrUpdateResponse], error)`
- New function `*FleetsClient.BeginDelete(context.Context, string, string, *FleetsClientBeginDeleteOptions) (*runtime.Poller[FleetsClientDeleteResponse], error)`
- New function `*FleetsClient.Get(context.Context, string, string, *FleetsClientGetOptions) (FleetsClientGetResponse, error)`
- New function `*FleetsClient.NewListByResourceGroupPager(string, *FleetsClientListByResourceGroupOptions) *runtime.Pager[FleetsClientListByResourceGroupResponse]`
- New function `*FleetsClient.ListCredentials(context.Context, string, string, *FleetsClientListCredentialsOptions) (FleetsClientListCredentialsResponse, error)`
- New function `*FleetsClient.NewListPager(*FleetsClientListOptions) *runtime.Pager[FleetsClientListResponse]`
- New function `*FleetsClient.Update(context.Context, string, string, *FleetsClientUpdateOptions) (FleetsClientUpdateResponse, error)`
- New function `NewManagedClusterSnapshotsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedClusterSnapshotsClient, error)`
- New function `*ManagedClusterSnapshotsClient.CreateOrUpdate(context.Context, string, string, ManagedClusterSnapshot, *ManagedClusterSnapshotsClientCreateOrUpdateOptions) (ManagedClusterSnapshotsClientCreateOrUpdateResponse, error)`
- New function `*ManagedClusterSnapshotsClient.Delete(context.Context, string, string, *ManagedClusterSnapshotsClientDeleteOptions) (ManagedClusterSnapshotsClientDeleteResponse, error)`
- New function `*ManagedClusterSnapshotsClient.Get(context.Context, string, string, *ManagedClusterSnapshotsClientGetOptions) (ManagedClusterSnapshotsClientGetResponse, error)`
- New function `*ManagedClusterSnapshotsClient.NewListByResourceGroupPager(string, *ManagedClusterSnapshotsClientListByResourceGroupOptions) *runtime.Pager[ManagedClusterSnapshotsClientListByResourceGroupResponse]`
- New function `*ManagedClusterSnapshotsClient.NewListPager(*ManagedClusterSnapshotsClientListOptions) *runtime.Pager[ManagedClusterSnapshotsClientListResponse]`
- New function `*ManagedClusterSnapshotsClient.UpdateTags(context.Context, string, string, TagsObject, *ManagedClusterSnapshotsClientUpdateTagsOptions) (ManagedClusterSnapshotsClientUpdateTagsResponse, error)`
- New function `*ManagedClustersClient.BeginAbortLatestOperation(context.Context, string, string, *ManagedClustersClientBeginAbortLatestOperationOptions) (*runtime.Poller[ManagedClustersClientAbortLatestOperationResponse], error)`
- New function `NewTrustedAccessRoleBindingsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*TrustedAccessRoleBindingsClient, error)`
- New function `*TrustedAccessRoleBindingsClient.CreateOrUpdate(context.Context, string, string, string, TrustedAccessRoleBinding, *TrustedAccessRoleBindingsClientCreateOrUpdateOptions) (TrustedAccessRoleBindingsClientCreateOrUpdateResponse, error)`
- New function `*TrustedAccessRoleBindingsClient.Delete(context.Context, string, string, string, *TrustedAccessRoleBindingsClientDeleteOptions) (TrustedAccessRoleBindingsClientDeleteResponse, error)`
- New function `*TrustedAccessRoleBindingsClient.Get(context.Context, string, string, string, *TrustedAccessRoleBindingsClientGetOptions) (TrustedAccessRoleBindingsClientGetResponse, error)`
- New function `*TrustedAccessRoleBindingsClient.NewListPager(string, string, *TrustedAccessRoleBindingsClientListOptions) *runtime.Pager[TrustedAccessRoleBindingsClientListResponse]`
- New function `NewTrustedAccessRolesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*TrustedAccessRolesClient, error)`
- New function `*TrustedAccessRolesClient.NewListPager(string, *TrustedAccessRolesClientListOptions) *runtime.Pager[TrustedAccessRolesClientListResponse]`
- New struct `AbsoluteMonthlySchedule`
- New struct `AgentPoolNetworkProfile`
- New struct `AgentPoolWindowsProfile`
- New struct `AgentPoolsClientAbortLatestOperationResponse`
- New struct `DailySchedule`
- New struct `DateSpan`
- New struct `Fleet`
- New struct `FleetCredentialResult`
- New struct `FleetCredentialResults`
- New struct `FleetHubProfile`
- New struct `FleetListResult`
- New struct `FleetMember`
- New struct `FleetMemberProperties`
- New struct `FleetMembersClient`
- New struct `FleetMembersClientCreateOrUpdateResponse`
- New struct `FleetMembersClientDeleteResponse`
- New struct `FleetMembersClientListByFleetResponse`
- New struct `FleetMembersListResult`
- New struct `FleetPatch`
- New struct `FleetProperties`
- New struct `FleetsClient`
- New struct `FleetsClientCreateOrUpdateResponse`
- New struct `FleetsClientDeleteResponse`
- New struct `FleetsClientListByResourceGroupResponse`
- New struct `FleetsClientListResponse`
- New struct `GuardrailsProfile`
- New struct `IPTag`
- New struct `MaintenanceWindow`
- New struct `ManagedClusterAzureMonitorProfile`
- New struct `ManagedClusterAzureMonitorProfileKubeStateMetrics`
- New struct `ManagedClusterAzureMonitorProfileMetrics`
- New struct `ManagedClusterIngressProfile`
- New struct `ManagedClusterIngressProfileWebAppRouting`
- New struct `ManagedClusterNodeResourceGroupProfile`
- New struct `ManagedClusterPropertiesForSnapshot`
- New struct `ManagedClusterSecurityProfileImageCleaner`
- New struct `ManagedClusterSecurityProfileNodeRestriction`
- New struct `ManagedClusterSecurityProfileWorkloadIdentity`
- New struct `ManagedClusterSnapshot`
- New struct `ManagedClusterSnapshotListResult`
- New struct `ManagedClusterSnapshotProperties`
- New struct `ManagedClusterSnapshotsClient`
- New struct `ManagedClusterSnapshotsClientListByResourceGroupResponse`
- New struct `ManagedClusterSnapshotsClientListResponse`
- New struct `ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler`
- New struct `ManagedClustersClientAbortLatestOperationResponse`
- New struct `NetworkProfileForSnapshot`
- New struct `NetworkProfileKubeProxyConfig`
- New struct `NetworkProfileKubeProxyConfigIpvsConfig`
- New struct `PortRange`
- New struct `RelativeMonthlySchedule`
- New struct `Schedule`
- New struct `TrustedAccessRole`
- New struct `TrustedAccessRoleBinding`
- New struct `TrustedAccessRoleBindingListResult`
- New struct `TrustedAccessRoleBindingProperties`
- New struct `TrustedAccessRoleBindingsClient`
- New struct `TrustedAccessRoleBindingsClientListResponse`
- New struct `TrustedAccessRoleListResult`
- New struct `TrustedAccessRoleRule`
- New struct `TrustedAccessRolesClient`
- New struct `TrustedAccessRolesClientListResponse`
- New struct `WeeklySchedule`
- New field `IgnorePodDisruptionBudget` in struct `AgentPoolsClientBeginDeleteOptions`
- New field `MaintenanceWindow` in struct `MaintenanceConfigurationProperties`
- New field `EnableVnetIntegration` in struct `ManagedClusterAPIServerAccessProfile`
- New field `SubnetID` in struct `ManagedClusterAPIServerAccessProfile`
- New field `CapacityReservationGroupID` in struct `ManagedClusterAgentPoolProfile`
- New field `EnableCustomCATrust` in struct `ManagedClusterAgentPoolProfile`
- New field `MessageOfTheDay` in struct `ManagedClusterAgentPoolProfile`
- New field `NetworkProfile` in struct `ManagedClusterAgentPoolProfile`
- New field `WindowsProfile` in struct `ManagedClusterAgentPoolProfile`
- New field `CapacityReservationGroupID` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `EnableCustomCATrust` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `MessageOfTheDay` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `NetworkProfile` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `WindowsProfile` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `NodeOSUpgradeChannel` in struct `ManagedClusterAutoUpgradeProfile`
- New field `EffectiveNoProxy` in struct `ManagedClusterHTTPProxyConfig`
- New field `BackendPoolType` in struct `ManagedClusterLoadBalancerProfile`
- New field `AzureMonitorProfile` in struct `ManagedClusterProperties`
- New field `CreationData` in struct `ManagedClusterProperties`
- New field `EnableNamespaceResources` in struct `ManagedClusterProperties`
- New field `GuardrailsProfile` in struct `ManagedClusterProperties`
- New field `IngressProfile` in struct `ManagedClusterProperties`
- New field `NodeResourceGroupProfile` in struct `ManagedClusterProperties`
- New field `CustomCATrustCertificates` in struct `ManagedClusterSecurityProfile`
- New field `ImageCleaner` in struct `ManagedClusterSecurityProfile`
- New field `NodeRestriction` in struct `ManagedClusterSecurityProfile`
- New field `WorkloadIdentity` in struct `ManagedClusterSecurityProfile`
- New field `Version` in struct `ManagedClusterStorageProfileDiskCSIDriver`
- New field `VerticalPodAutoscaler` in struct `ManagedClusterWorkloadAutoScalerProfile`
- New field `IgnorePodDisruptionBudget` in struct `ManagedClustersClientBeginDeleteOptions`
- New field `EbpfDataplane` in struct `NetworkProfile`
- New field `KubeProxyConfig` in struct `NetworkProfile`
- New field `NetworkPluginMode` in struct `NetworkProfile`


## 2.3.0 (2023-01-27)
### Features Added

- New value `ManagedClusterPodIdentityProvisioningStateCanceled`, `ManagedClusterPodIdentityProvisioningStateSucceeded` added to type alias `ManagedClusterPodIdentityProvisioningState`
- New value `PrivateEndpointConnectionProvisioningStateCanceled` added to type alias `PrivateEndpointConnectionProvisioningState`
- New struct `ManagedClusterWorkloadAutoScalerProfile`
- New struct `ManagedClusterWorkloadAutoScalerProfileKeda`
- New field `WorkloadAutoScalerProfile` in struct `ManagedClusterProperties`
- New field `Location` in struct `ManagedClustersClientGetCommandResultResponse`


## 2.2.0 (2022-10-26)
### Features Added

- New function `*ManagedClustersClient.BeginRotateServiceAccountSigningKeys(context.Context, string, string, *ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions) (*runtime.Poller[ManagedClustersClientRotateServiceAccountSigningKeysResponse], error)`
- New struct `ManagedClusterOIDCIssuerProfile`
- New struct `ManagedClusterStorageProfileBlobCSIDriver`
- New struct `ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions`
- New struct `ManagedClustersClientRotateServiceAccountSigningKeysResponse`
- New field `BlobCSIDriver` in struct `ManagedClusterStorageProfile`
- New field `OidcIssuerProfile` in struct `ManagedClusterProperties`


## 2.1.0 (2022-08-25)
### Features Added

- New const `OSSKUWindows2019`
- New const `OSSKUWindows2022`


## 2.0.0 (2022-07-22)
### Breaking Changes

- Struct `ManagedClusterSecurityProfileAzureDefender` has been removed
- Field `AzureDefender` of struct `ManagedClusterSecurityProfile` has been removed

### Features Added

- New const `KeyVaultNetworkAccessTypesPrivate`
- New const `NetworkPluginNone`
- New const `KeyVaultNetworkAccessTypesPublic`
- New function `PossibleKeyVaultNetworkAccessTypesValues() []KeyVaultNetworkAccessTypes`
- New struct `AzureKeyVaultKms`
- New struct `ManagedClusterSecurityProfileDefender`
- New struct `ManagedClusterSecurityProfileDefenderSecurityMonitoring`
- New field `HostGroupID` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `HostGroupID` in struct `ManagedClusterAgentPoolProfile`
- New field `AzureKeyVaultKms` in struct `ManagedClusterSecurityProfile`
- New field `Defender` in struct `ManagedClusterSecurityProfile`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
