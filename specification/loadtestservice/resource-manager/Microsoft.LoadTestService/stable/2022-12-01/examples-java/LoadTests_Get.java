
/**
 * Samples for LoadTests GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_Get
     * .json
     */
    /**
     * Sample code: Get a LoadTestResource.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void getALoadTestResource(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.loadTests().getByResourceGroupWithResponse("dummyrg", "myLoadTest", com.azure.core.util.Context.NONE);
    }
}
