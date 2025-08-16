
import com.azure.resourcemanager.dataprotection.models.AksVolumeTypes;
import com.azure.resourcemanager.dataprotection.models.AzureOperationalStoreParameters;
import com.azure.resourcemanager.dataprotection.models.BackupInstance;
import com.azure.resourcemanager.dataprotection.models.DataStoreTypes;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.KubernetesClusterBackupDatasourceParameters;
import com.azure.resourcemanager.dataprotection.models.PolicyInfo;
import com.azure.resourcemanager.dataprotection.models.PolicyParameters;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupInstances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/
     * BackupInstanceOperations/PutBackupInstance_KubernetesClusterBackupDatasourceParameters.json
     */
    /**
     * Sample code: Create BackupInstance With KubernetesClusterBackupDatasourceParameters.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void createBackupInstanceWithKubernetesClusterBackupDatasourceParameters(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().define("aksbi").withExistingBackupVault("aksrg", "aksvault")
            .withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withProperties(new BackupInstance().withFriendlyName("aksbi").withDataSourceInfo(new Datasource()
                .withDatasourceType("Microsoft.ContainerService/managedclusters").withObjectType("Datasource")
                .withResourceId(
                    "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster")
                .withResourceLocation("eastus2euap").withResourceName("akscluster")
                .withResourceType("Microsoft.ContainerService/managedclusters").withResourceUri(
                    "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"))
                .withDataSourceSetInfo(new DatasourceSet()
                    .withDatasourceType("Microsoft.ContainerService/managedclusters").withObjectType("DatasourceSet")
                    .withResourceId(
                        "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster")
                    .withResourceLocation("eastus2euap").withResourceName("akscluster")
                    .withResourceType("Microsoft.ContainerService/managedclusters").withResourceUri(
                        "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg/providers/Microsoft.ContainerService/managedClusters/akscluster"))
                .withPolicyInfo(new PolicyInfo().withPolicyId(
                    "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourcegroups/aksrg/providers/Microsoft.DataProtection/BackupVaults/aksvault/backupPolicies/akspolicy")
                    .withPolicyParameters(new PolicyParameters()
                        .withDataStoreParametersList(Arrays.asList(new AzureOperationalStoreParameters()
                            .withDataStoreType(DataStoreTypes.OPERATIONAL_STORE).withResourceGroupId(
                                "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/aksrg")))
                        .withBackupDatasourceParametersList(
                            Arrays.asList(new KubernetesClusterBackupDatasourceParameters().withSnapshotVolumes(true)
                                .withIncludedVolumeTypes(
                                    Arrays.asList(AksVolumeTypes.AZURE_DISK, AksVolumeTypes.AZURE_FILE_SHARE_SMB))
                                .withIncludeClusterScopeResources(true).withIncludedNamespaces(Arrays.asList("test"))
                                .withExcludedNamespaces(Arrays.asList("kube-system"))
                                .withIncludedResourceTypes(Arrays.asList())
                                .withExcludedResourceTypes(Arrays.asList("v1/Secret"))
                                .withLabelSelectors(Arrays.asList())))))
                .withObjectType("BackupInstance"))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
