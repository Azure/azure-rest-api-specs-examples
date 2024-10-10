
/**
 * Samples for Labs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/deleteLab.json
     */
    /**
     * Sample code: deleteLab.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void deleteLab(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labs().delete("testrg123", "testlab", com.azure.core.util.Context.NONE);
    }
}
