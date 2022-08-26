import com.azure.core.util.Context;

/** Samples for Usages ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Usages/getUsages.json
     */
    /**
     * Sample code: listUsages.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listUsages(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.usages().listByLocation("eastus2", null, Context.NONE);
    }
}
