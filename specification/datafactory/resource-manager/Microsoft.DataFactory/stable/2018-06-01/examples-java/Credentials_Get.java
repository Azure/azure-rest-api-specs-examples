import com.azure.core.util.Context;

/** Samples for CredentialOperations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_Get.json
     */
    /**
     * Sample code: Credentials_Get.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void credentialsGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .credentialOperations()
            .getWithResponse("exampleResourceGroup", "exampleFactoryName", "exampleCredential", null, Context.NONE);
    }
}
