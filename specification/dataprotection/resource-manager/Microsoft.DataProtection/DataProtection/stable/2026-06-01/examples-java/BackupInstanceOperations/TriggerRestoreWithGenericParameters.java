
import com.azure.resourcemanager.dataprotection.models.AzureBackupRecoveryPointBasedRestoreRequest;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.GenericRestoreDatasourceCriteria;
import com.azure.resourcemanager.dataprotection.models.ItemLevelRestoreTargetInfo;
import com.azure.resourcemanager.dataprotection.models.RecoveryOption;
import com.azure.resourcemanager.dataprotection.models.ResourceListSelectionCriteria;
import com.azure.resourcemanager.dataprotection.models.SourceDataStoreType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupInstances TriggerRestore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/BackupInstanceOperations/TriggerRestoreWithGenericParameters.json
     */
    /**
     * Sample code: Trigger Restore with GenericRestoreDatasourceCriteria.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void triggerRestoreWithGenericRestoreDatasourceCriteria(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances()
            .triggerRestore("000pikumar", "PrivatePreviewVault1", "testInstance1",
                new AzureBackupRecoveryPointBasedRestoreRequest()
                    .withRestoreTargetInfo(new ItemLevelRestoreTargetInfo()
                        .withRecoveryOption(RecoveryOption.FAIL_IF_EXISTS).withRestoreLocation("southeastasia")
                        .withRestoreCriteria(
                            Arrays.asList(new GenericRestoreDatasourceCriteria().withResourceSelectors(
                                new ResourceListSelectionCriteria().withObjectType("resourceListSelectionCriteria")
                                    .withResourceIdentifiers(Arrays
                                        .asList("source-vol1", "source-vol2", "source-vol3"))
                                    .withResourceNameOverrides(
                                        mapOf("source-vol1", "target-vol1", "source-vol2", "target-vol2")))))
                        .withDatasourceInfo(new Datasource()
                            .withDatasourceType("Microsoft.ElasticSan/elasticSans/volumeGroups")
                            .withObjectType("Datasource")
                            .withResourceId(
                                "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc/volumeGroups/target-esan-volgroup")
                            .withResourceLocation("eastus2euap").withResourceName("target-esan-volgroup")
                            .withResourceType(
                                "Microsoft.ElasticSan/elasticSans/volumeGroups")
                            .withResourceUri("SampleresourceUri123"))
                        .withDatasourceSetInfo(new DatasourceSet()
                            .withDatasourceType("Microsoft.ElasticSan/elasticSans/volumeGroups")
                            .withObjectType("DatasourceSet")
                            .withResourceId(
                                "/subscriptions/97cda027-4279-4cde-b4ff-19afa0021d87/resourceGroups/ESAN-ECYBVTRG/providers/Microsoft.ElasticSan/elasticSans/ecy-bvt-adhoc")
                            .withResourceLocation("eastus2euap").withResourceType("Microsoft.ElasticSan/elasticSans")))
                    .withSourceDataStoreType(SourceDataStoreType.OPERATIONAL_STORE).withRecoveryPointId("hardcodedRP"),
                com.azure.core.util.Context.NONE);
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
