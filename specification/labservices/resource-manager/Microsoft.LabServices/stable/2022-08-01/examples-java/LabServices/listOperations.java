import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabServices/listOperations.json
     */
    /**
     * Sample code: listOperations.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listOperations(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.operations().list(Context.NONE);
    }
}
