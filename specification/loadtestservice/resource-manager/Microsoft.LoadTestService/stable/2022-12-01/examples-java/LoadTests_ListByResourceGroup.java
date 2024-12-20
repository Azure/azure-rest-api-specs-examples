
/**
 * Samples for LoadTests ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/
     * LoadTests_ListByResourceGroup.json
     */
    /**
     * Sample code: List LoadTestResource resources by resource group.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void
        listLoadTestResourceResourcesByResourceGroup(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.loadTests().listByResourceGroup("dummyrg", com.azure.core.util.Context.NONE);
    }
}
