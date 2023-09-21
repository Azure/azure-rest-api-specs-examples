/** Samples for Certificates GenerateVerificationCode. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_generateverificationcode.json
     */
    /**
     * Sample code: Certificates_GenerateVerificationCode.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesGenerateVerificationCode(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .certificates()
            .generateVerificationCodeWithResponse(
                "myResourceGroup", "testHub", "cert", "AAAAAAAADGk=", com.azure.core.util.Context.NONE);
    }
}
