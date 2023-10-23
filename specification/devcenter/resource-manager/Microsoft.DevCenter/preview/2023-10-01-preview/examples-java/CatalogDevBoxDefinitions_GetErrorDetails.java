/** Samples for CatalogDevBoxDefinitions GetErrorDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/CatalogDevBoxDefinitions_GetErrorDetails.json
     */
    /**
     * Sample code: CatalogDevBoxDefinitions_GetErrorDetails.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogDevBoxDefinitionsGetErrorDetails(
        com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .catalogDevBoxDefinitions()
            .getErrorDetailsWithResponse(
                "rg1", "Contoso", "CentralCatalog", "WebDevBox", com.azure.core.util.Context.NONE);
    }
}
