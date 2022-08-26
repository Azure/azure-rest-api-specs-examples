import com.azure.core.util.Context;

/** Samples for Labs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/listLabs.json
     */
    /**
     * Sample code: listLabs.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listLabs(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labs().list(null, Context.NONE);
    }
}
