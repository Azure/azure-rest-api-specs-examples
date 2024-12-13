
import com.azure.resourcemanager.loadtesting.models.LoadTestResource;

/**
 * Samples for LoadTests Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/
     * LoadTests_Update.json
     */
    /**
     * Sample code: Update a LoadTestResource.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void updateALoadTestResource(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        LoadTestResource resource = manager.loadTests()
            .getByResourceGroupWithResponse("dummyrg", "myLoadTest", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
