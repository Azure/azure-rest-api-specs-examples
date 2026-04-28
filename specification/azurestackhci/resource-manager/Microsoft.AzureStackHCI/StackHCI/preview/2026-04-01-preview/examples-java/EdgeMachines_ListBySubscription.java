
/**
 * Samples for EdgeMachines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachines_ListBySubscription.json
     */
    /**
     * Sample code: List edge machines in a given subscription.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listEdgeMachinesInAGivenSubscription(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachines().list(com.azure.core.util.Context.NONE);
    }
}
