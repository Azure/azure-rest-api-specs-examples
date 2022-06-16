import com.azure.core.util.Context;

/** Samples for Images Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Images/getImage.json
     */
    /**
     * Sample code: getImage.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void getImage(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.images().getWithResponse("testrg123", "testlabplan", "image1", Context.NONE);
    }
}
