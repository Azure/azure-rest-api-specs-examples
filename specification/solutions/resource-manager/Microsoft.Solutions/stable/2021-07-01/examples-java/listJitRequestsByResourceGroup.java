/** Samples for JitRequests ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listJitRequestsByResourceGroup.json
     */
    /**
     * Sample code: Lists all JIT requests within the resource group.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listsAllJITRequestsWithinTheResourceGroup(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.jitRequests().listByResourceGroupWithResponse("rg", com.azure.core.util.Context.NONE);
    }
}
