import com.azure.core.util.Context;

/** Samples for Labs Publish. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/publishLab.json
     */
    /**
     * Sample code: publishLab.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void publishLab(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labs().publish("testrg123", "testlab", Context.NONE);
    }
}
