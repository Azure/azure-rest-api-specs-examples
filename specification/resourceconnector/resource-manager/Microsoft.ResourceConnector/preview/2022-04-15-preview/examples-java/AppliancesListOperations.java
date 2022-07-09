import com.azure.core.util.Context;

/** Samples for Appliances ListOperations. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesListOperations.json
     */
    /**
     * Sample code: List Appliances operations.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void listAppliancesOperations(com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager.appliances().listOperations(Context.NONE);
    }
}
