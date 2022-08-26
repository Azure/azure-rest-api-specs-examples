import com.azure.core.util.Context;

/** Samples for LabPlans List. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/listLabPlans.json
     */
    /**
     * Sample code: listLabPlans.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listLabPlans(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labPlans().list(null, Context.NONE);
    }
}
