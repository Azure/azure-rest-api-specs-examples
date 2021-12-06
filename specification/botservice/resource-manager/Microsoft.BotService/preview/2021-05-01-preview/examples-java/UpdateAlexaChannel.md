Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.2/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.botservice.models.AlexaChannel;
import com.azure.resourcemanager.botservice.models.AlexaChannelProperties;
import com.azure.resourcemanager.botservice.models.BotChannel;

/** Samples for Channels Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateAlexaChannel.json
     */
    /**
     * Sample code: Update Alexa.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void updateAlexa(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        BotChannel resource =
            manager
                .channels()
                .getWithResponse("OneResourceGroupName", "samplebotname", "AlexaChannel", Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(
                new AlexaChannel()
                    .withProperties(
                        new AlexaChannelProperties().withAlexaSkillId("XAlexaSkillIdX").withIsEnabled(true)))
            .apply();
    }
}
```
