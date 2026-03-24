
/**
 * Samples for CredentialOperations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Credentials_Delete.json
     */
    /**
     * Sample code: Credentials_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void credentialsDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.credentialOperations().deleteWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleCredential", com.azure.core.util.Context.NONE);
    }
}
