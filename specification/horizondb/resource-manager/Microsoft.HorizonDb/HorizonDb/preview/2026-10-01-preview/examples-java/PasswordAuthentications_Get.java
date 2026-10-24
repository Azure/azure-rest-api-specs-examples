
/**
 * Samples for PasswordAuthentications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PasswordAuthentications_Get.json
     */
    /**
     * Sample code: Get the password authentication method.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getThePasswordAuthenticationMethod(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.passwordAuthentications().getWithResponse("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
