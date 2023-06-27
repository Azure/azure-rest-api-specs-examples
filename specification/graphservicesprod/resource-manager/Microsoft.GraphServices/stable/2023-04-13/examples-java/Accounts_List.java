/** Samples for Accounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Accounts_List.json
     */
    /**
     * Sample code: Create or update account resource.
     *
     * @param manager Entry point to GraphServicesManager.
     */
    public static void createOrUpdateAccountResource(
        com.azure.resourcemanager.graphservices.GraphServicesManager manager) {
        manager.accounts().listByResourceGroup("testResourceGroupGRAM", com.azure.core.util.Context.NONE);
    }
}
