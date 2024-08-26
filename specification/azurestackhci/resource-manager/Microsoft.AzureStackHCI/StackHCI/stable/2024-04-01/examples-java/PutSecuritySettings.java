
import com.azure.resourcemanager.azurestackhci.models.ComplianceAssignmentType;

/**
 * Samples for SecuritySettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * PutSecuritySettings.json
     */
    /**
     * Sample code: Create Security Settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createSecuritySettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.securitySettings().define("default").withExistingCluster("test-rg", "myCluster")
            .withSecuredCoreComplianceAssignment(ComplianceAssignmentType.AUDIT)
            .withWdacComplianceAssignment(ComplianceAssignmentType.APPLY_AND_AUTO_CORRECT)
            .withSmbEncryptionForIntraClusterTrafficComplianceAssignment(ComplianceAssignmentType.AUDIT).create();
    }
}
