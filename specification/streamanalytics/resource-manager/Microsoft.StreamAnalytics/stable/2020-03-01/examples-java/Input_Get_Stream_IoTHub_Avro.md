Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Inputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Get_Stream_IoTHub_Avro.json
     */
    /**
     * Sample code: Get a stream IoT Hub input with Avro serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAStreamIoTHubInputWithAvroSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().getWithResponse("sjrg3467", "sj9742", "input7970", Context.NONE);
    }
}
```
