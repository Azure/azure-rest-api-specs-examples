```java
import com.azure.core.util.Context;

/** Samples for Organization Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Delete.json
     */
    /**
     * Sample code: Confluent_Delete.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void confluentDelete(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().delete("myResourceGroup", "myOrganization", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-confluent_1.0.0-beta.3/sdk/confluent/azure-resourcemanager-confluent/README.md) on how to add the SDK to your project and authenticate.
