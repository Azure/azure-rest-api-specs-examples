
/**
 * Samples for ProviderActions ListInUseStorageAccounts.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/ListInUseStorageAccountsWithSubscriptions.json
     */
    /**
     * Sample code: ListInUseStorageAccountsWithSubscriptions.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void listInUseStorageAccountsWithSubscriptions(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.providerActions().listInUseStorageAccountsWithResponse(null, com.azure.core.util.Context.NONE);
    }
}
