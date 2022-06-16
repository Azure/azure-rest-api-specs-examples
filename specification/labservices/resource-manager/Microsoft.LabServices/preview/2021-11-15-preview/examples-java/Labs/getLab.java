import com.azure.core.util.Context;

/** Samples for Labs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/getLab.json
     */
    /**
     * Sample code: getLab.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void getLab(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labs().getByResourceGroupWithResponse("testrg123", "testlab", Context.NONE);
    }
}
