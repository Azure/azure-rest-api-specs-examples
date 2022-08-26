import com.azure.resourcemanager.labservices.models.EnableState;

/** Samples for Images CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Images/putImage.json
     */
    /**
     * Sample code: putImage.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void putImage(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .images()
            .define("image1")
            .withExistingLabPlan("testrg123", "testlabplan")
            .withEnabledState(EnableState.ENABLED)
            .create();
    }
}
