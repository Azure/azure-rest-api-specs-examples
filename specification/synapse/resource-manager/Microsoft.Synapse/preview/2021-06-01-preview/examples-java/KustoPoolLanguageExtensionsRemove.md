Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.LanguageExtensionInner;
import com.azure.resourcemanager.synapse.models.LanguageExtensionName;
import com.azure.resourcemanager.synapse.models.LanguageExtensionsList;
import java.util.Arrays;

/** Samples for KustoPools RemoveLanguageExtensions. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolLanguageExtensionsRemove.json
     */
    /**
     * Sample code: KustoPoolRemoveLanguageExtensions.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolRemoveLanguageExtensions(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .removeLanguageExtensions(
                "kustorptest",
                "kustoclusterrptest4",
                "kustorptest",
                new LanguageExtensionsList()
                    .withValue(
                        Arrays
                            .asList(
                                new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.PYTHON),
                                new LanguageExtensionInner().withLanguageExtensionName(LanguageExtensionName.R))),
                Context.NONE);
    }
}
```
