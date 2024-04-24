
/**
 * Samples for EnvironmentDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * EnvironmentDefinitions_Get.json
     */
    /**
     * Sample code: EnvironmentDefinitions_Get.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void environmentDefinitionsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentDefinitions().getWithResponse("rg1", "Contoso", "myCatalog", "myEnvironmentDefinition",
            com.azure.core.util.Context.NONE);
    }
}
