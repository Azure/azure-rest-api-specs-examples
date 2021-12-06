Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-botservice_1.0.0-beta.2/sdk/botservice/azure-resourcemanager-botservice/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.botservice.models.ConnectionSettingParameter;
import com.azure.resourcemanager.botservice.models.ConnectionSettingProperties;
import java.util.Arrays;

/** Samples for BotConnection Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutConnection.json
     */
    /**
     * Sample code: Create Connection Setting.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void createConnectionSetting(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .botConnections()
            .define("sampleConnection")
            .withRegion("West US")
            .withExistingBotService("OneResourceGroupName", "samplebotname")
            .withProperties(
                new ConnectionSettingProperties()
                    .withClientId("sampleclientid")
                    .withClientSecret("samplesecret")
                    .withScopes("samplescope")
                    .withServiceProviderId("serviceproviderid")
                    .withParameters(
                        Arrays
                            .asList(
                                new ConnectionSettingParameter().withKey("key1").withValue("value1"),
                                new ConnectionSettingParameter().withKey("key2").withValue("value2"))))
            .withEtag("etag1")
            .create();
    }
}
```
