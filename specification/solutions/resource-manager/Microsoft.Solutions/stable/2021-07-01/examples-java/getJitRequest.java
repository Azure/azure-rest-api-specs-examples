/** Samples for JitRequests GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/getJitRequest.json
     */
    /**
     * Sample code: Gets the jit request.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void getsTheJitRequest(com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.jitRequests().getByResourceGroupWithResponse("rg", "myJitRequest", com.azure.core.util.Context.NONE);
    }
}
