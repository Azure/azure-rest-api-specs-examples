import com.azure.core.util.Context;

/** Samples for Usages ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Usages/getUsages.json
     */
    /**
     * Sample code: listUsages.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listUsages(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.usages().listByLocation("westus2", null, Context.NONE);
    }
}
