
/**
 * Samples for EdgeMachines ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachines_ListByResourceGroup.json
     */
    /**
     * Sample code: List edge machines in a given resource group.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listEdgeMachinesInAGivenResourceGroup(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachines().listByResourceGroup("ArcInstance-rg", com.azure.core.util.Context.NONE);
    }
}
