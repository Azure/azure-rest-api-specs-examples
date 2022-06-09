```java
/** Samples for SignalRCustomCertificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomCertificates_CreateOrUpdate.json
     */
    /**
     * Sample code: SignalRCustomCertificates_CreateOrUpdate.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomCertificatesCreateOrUpdate(
        com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRCustomCertificates()
            .define("myCert")
            .withExistingSignalR("myResourceGroup", "mySignalRService")
            .withKeyVaultBaseUri("https://myvault.keyvault.azure.net/")
            .withKeyVaultSecretName("mycert")
            .withKeyVaultSecretVersion("bb6a44b2743f47f68dad0d6cc9756432")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-signalr_1.0.0-beta.4/sdk/signalr/azure-resourcemanager-signalr/README.md) on how to add the SDK to your project and authenticate.
