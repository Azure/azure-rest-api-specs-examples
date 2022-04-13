Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-cognitiveservices_1.0.0-beta.4/sdk/cognitiveservices/azure-resourcemanager-cognitiveservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Accounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/ListAccountsByResourceGroup.json
     */
    /**
     * Sample code: List Accounts by Resource Group.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listAccountsByResourceGroup(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
```
