```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cognitiveservices.models.KeyName;
import com.azure.resourcemanager.cognitiveservices.models.RegenerateKeyParameters;

/** Samples for Accounts RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/RegenerateKey.json
     */
    /**
     * Sample code: Regenerate Keys.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void regenerateKeys(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .accounts()
            .regenerateKeyWithResponse(
                "myResourceGroup", "myAccount", new RegenerateKeyParameters().withKeyName(KeyName.KEY2), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-cognitiveservices_1.0.0-beta.4/sdk/cognitiveservices/azure-resourcemanager-cognitiveservices/README.md) on how to add the SDK to your project and authenticate.
