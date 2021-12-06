Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-logz_1.0.0-beta.1/sdk/logz/azure-resourcemanager-logz/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for SingleSignOn CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SingleSignOnConfigurations_CreateOrUpdate.json
     */
    /**
     * Sample code: SingleSignOnConfigurations_CreateOrUpdate.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void singleSignOnConfigurationsCreateOrUpdate(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.singleSignOns().define("default").withExistingMonitor("myResourceGroup", "myMonitor").create();
    }
}
```
