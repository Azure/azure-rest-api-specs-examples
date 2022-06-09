```java
/** Samples for DicomServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Create.json
     */
    /**
     * Sample code: Create or update a Dicom Service.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void createOrUpdateADicomService(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager
            .dicomServices()
            .define("blue")
            .withExistingWorkspace("testRG", "workspace1")
            .withRegion("westus")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-healthcareapis_1.0.0-beta.2/sdk/healthcareapis/azure-resourcemanager-healthcareapis/README.md) on how to add the SDK to your project and authenticate.
