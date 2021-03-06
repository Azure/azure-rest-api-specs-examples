import com.azure.core.util.Context;

/** Samples for LoadTests ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_ListByResourceGroup.json
     */
    /**
     * Sample code: LoadTests_ListByResourceGroup.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void loadTestsListByResourceGroup(com.azure.resourcemanager.loadtestservice.LoadTestManager manager) {
        manager.loadTests().listByResourceGroup("dummyrg", Context.NONE);
    }
}
