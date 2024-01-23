
/**
 * Samples for CredentialOperations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_Delete.
     * json
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
