
/**
 * Samples for LoadTests CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/
     * LoadTests_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a LoadTestResource.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void createALoadTestResource(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.loadTests().define("myLoadTest").withRegion((String) null).withExistingResourceGroup("dummyrg")
            .create();
    }
}
