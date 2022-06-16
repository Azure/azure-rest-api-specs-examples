import com.azure.core.util.Context;

/** Samples for Images ListByLabPlan. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Images/listImages.json
     */
    /**
     * Sample code: listImages.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listImages(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.images().listByLabPlan("testrg123", "testlabplan", null, Context.NONE);
    }
}
