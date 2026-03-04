
/**
 * Samples for AzureTrafficCollectorsByResourceGroup ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2022-11-01/AzureTrafficCollectorsByResourceGroupList.json
     */
    /**
     * Sample code: List of Traffic Collectors by ResourceGroup.
     * 
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void listOfTrafficCollectorsByResourceGroup(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.azureTrafficCollectorsByResourceGroups().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
