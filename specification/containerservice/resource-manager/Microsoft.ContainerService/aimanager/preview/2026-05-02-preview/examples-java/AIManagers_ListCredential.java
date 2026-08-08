
/**
 * Samples for AIManagers ListCredential.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIManagers_ListCredential.json
     */
    /**
     * Sample code: Lists the credentials of an AI Manager.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void listsTheCredentialsOfAnAIManager(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagers().listCredentialWithResponse("rg1", "aimanager1", com.azure.core.util.Context.NONE);
    }
}
