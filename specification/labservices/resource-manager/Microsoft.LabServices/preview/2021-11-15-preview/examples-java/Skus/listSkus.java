import com.azure.core.util.Context;

/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Skus/listSkus.json
     */
    /**
     * Sample code: listSkus.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listSkus(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.skus().list(null, Context.NONE);
    }
}
