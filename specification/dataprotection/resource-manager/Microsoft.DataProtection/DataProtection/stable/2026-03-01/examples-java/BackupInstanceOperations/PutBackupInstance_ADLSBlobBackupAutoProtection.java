
import com.azure.resourcemanager.dataprotection.models.AdlsBlobBackupDatasourceParametersForAutoProtection;
import com.azure.resourcemanager.dataprotection.models.BackupInstance;
import com.azure.resourcemanager.dataprotection.models.BlobBackupAutoProtectionRule;
import com.azure.resourcemanager.dataprotection.models.BlobBackupPatternType;
import com.azure.resourcemanager.dataprotection.models.BlobBackupRuleBasedAutoProtectionSettings;
import com.azure.resourcemanager.dataprotection.models.BlobBackupRuleMode;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.PolicyInfo;
import com.azure.resourcemanager.dataprotection.models.PolicyParameters;
import java.util.Arrays;

/**
 * Samples for BackupInstances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/BackupInstanceOperations/PutBackupInstance_ADLSBlobBackupAutoProtection.json
     */
    /**
     * Sample code: Create BackupInstance With ADLSBlobBackupAutoProtection.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void createBackupInstanceWithADLSBlobBackupAutoProtection(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().define("adlsstorageaccount-adlsstorageaccount-3a76f8a-c176-4f7d-819e-95157e2b0071")
            .withExistingBackupVault("adlsrg", "adlsvault")
            .withProperties(new BackupInstance().withFriendlyName("adlsstorageaccount\\adlsbackupinstance")
                .withDataSourceInfo(new Datasource()
                    .withDatasourceType("Microsoft.Storage/storageAccounts/adlsBlobServices")
                    .withObjectType("Datasource")
                    .withResourceId(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount")
                    .withResourceLocation("centraluseuap").withResourceName("adlsstorageaccount")
                    .withResourceType("microsoft.storage/storageAccounts").withResourceUri(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount"))
                .withDataSourceSetInfo(new DatasourceSet()
                    .withDatasourceType("Microsoft.Storage/storageAccounts/adlsBlobServices")
                    .withObjectType("DatasourceSet")
                    .withResourceId(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount")
                    .withResourceLocation("centraluseuap").withResourceName("adlsstorageaccount")
                    .withResourceType("microsoft.storage/storageAccounts").withResourceUri(
                        "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.Storage/storageAccounts/adlsstorageaccount"))
                .withPolicyInfo(new PolicyInfo().withPolicyId(
                    "/subscriptions/54707983-993e-43de-8d94-074451394eda/resourceGroups/adlsrg/providers/Microsoft.DataProtection/backupVaults/adlsvault/backupPolicies/adlspolicy")
                    .withPolicyParameters(new PolicyParameters().withBackupDatasourceParametersList(Arrays
                        .asList(new AdlsBlobBackupDatasourceParametersForAutoProtection().withAutoProtectionSettings(
                            new BlobBackupRuleBasedAutoProtectionSettings().withEnabled(true)
                                .withRules(Arrays.asList(
                                    new BlobBackupAutoProtectionRule().withObjectType("BlobBackupAutoProtectionRule")
                                        .withMode(BlobBackupRuleMode.EXCLUDE).withType(BlobBackupPatternType.PREFIX)
                                        .withPattern("temp-"),
                                    new BlobBackupAutoProtectionRule().withObjectType("BlobBackupAutoProtectionRule")
                                        .withMode(BlobBackupRuleMode.EXCLUDE).withType(BlobBackupPatternType.PREFIX)
                                        .withPattern("test-"))))))))
                .withObjectType("BackupInstance"))
            .create();
    }
}
