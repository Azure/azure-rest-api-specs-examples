
/**
 * Samples for CredentialOperations ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * Credentials_ListByFactory.json
     */
    /**
     * Sample code: Credentials_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void credentialsListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.credentialOperations().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
