
/**
 * Samples for LoadTests Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/
     * LoadTests_Delete.json
     */
    /**
     * Sample code: Delete a LoadTestResource.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void deleteALoadTestResource(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.loadTests().delete("dummyrg", "myLoadTest", com.azure.core.util.Context.NONE);
    }
}
