
/**
 * Samples for ProviderActions ListInUseStorageAccounts.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/
     * examples/ListInUseStorageAccountsWithoutSubscriptions.json
     */
    /**
     * Sample code: ListInUseStorageAccountsWithoutSubscriptions.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void listInUseStorageAccountsWithoutSubscriptions(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.providerActions().listInUseStorageAccountsWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
