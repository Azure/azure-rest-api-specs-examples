Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.DeliveryRuleResponseHeaderAction;
import com.azure.resourcemanager.cdn.models.HeaderAction;
import com.azure.resourcemanager.cdn.models.HeaderActionParameters;
import com.azure.resourcemanager.cdn.models.RuleUpdateParameters;
import java.util.Arrays;

/** Samples for Rules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Update.json
     */
    /**
     * Sample code: Rules_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rulesUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRules()
            .update(
                "RG",
                "profile1",
                "ruleSet1",
                "rule1",
                new RuleUpdateParameters()
                    .withOrder(1)
                    .withActions(
                        Arrays
                            .asList(
                                new DeliveryRuleResponseHeaderAction()
                                    .withParameters(
                                        new HeaderActionParameters()
                                            .withHeaderAction(HeaderAction.OVERWRITE)
                                            .withHeaderName("X-CDN")
                                            .withValue("MSFT")))),
                Context.NONE);
    }
}
```
