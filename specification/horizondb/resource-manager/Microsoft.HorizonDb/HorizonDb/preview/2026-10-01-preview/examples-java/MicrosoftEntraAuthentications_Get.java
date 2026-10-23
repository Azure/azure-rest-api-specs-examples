
/**
 * Samples for MicrosoftEntraAuthentications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MicrosoftEntraAuthentications_Get.json
     */
    /**
     * Sample code: Get the Microsoft Entra authentication method.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getTheMicrosoftEntraAuthenticationMethod(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.microsoftEntraAuthentications().getWithResponse("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
