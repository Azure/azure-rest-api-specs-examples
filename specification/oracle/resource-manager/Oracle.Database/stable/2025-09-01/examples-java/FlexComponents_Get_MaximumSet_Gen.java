
/**
 * Samples for FlexComponents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/FlexComponents_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: FlexComponents_Get_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        flexComponentsGetMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.flexComponents().getWithResponse("eastus", "flexname1", com.azure.core.util.Context.NONE);
    }
}
