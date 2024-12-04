
/**
 * Samples for EnvironmentDefinitions GetErrorDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * EnvironmentDefinitions_GetErrorDetails.json
     */
    /**
     * Sample code: EnvironmentDefinitions_GetErrorDetails.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void
        environmentDefinitionsGetErrorDetails(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentDefinitions().getErrorDetailsWithResponse("rg1", "Contoso", "myCatalog",
            "myEnvironmentDefinition", com.azure.core.util.Context.NONE);
    }
}
