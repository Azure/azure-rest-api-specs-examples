
/**
 * Samples for Factories List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_List.json
     */
    /**
     * Sample code: Factories_List.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesList(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.factories().list(com.azure.core.util.Context.NONE);
    }
}
