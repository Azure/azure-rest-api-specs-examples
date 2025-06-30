
/**
 * Samples for AuthenticationSettings ListByHealthModel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/AuthenticationSettings_ListByHealthModel.json
     */
    /**
     * Sample code: AuthenticationSettings_ListByHealthModel.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        authenticationSettingsListByHealthModel(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.authenticationSettings().listByHealthModel("my-resource-group", "my-health-model",
            com.azure.core.util.Context.NONE);
    }
}
