Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.5.0/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_CreateOrUpdate.json
func ExampleIntegrationAccountAgreementsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlogic.NewIntegrationAccountAgreementsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<integration-account-name>",
		"<agreement-name>",
		armlogic.IntegrationAccountAgreement{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"IntegrationAccountAgreement": to.Ptr("<IntegrationAccountAgreementName>"),
			},
			Properties: &armlogic.IntegrationAccountAgreementProperties{
				AgreementType: to.Ptr(armlogic.AgreementTypeAS2),
				Content: &armlogic.AgreementContent{
					AS2: &armlogic.AS2AgreementContent{
						ReceiveAgreement: &armlogic.AS2OneWayAgreement{
							ProtocolSettings: &armlogic.AS2ProtocolSettings{
								AcknowledgementConnectionSettings: &armlogic.AS2AcknowledgementConnectionSettings{
									IgnoreCertificateNameMismatch: to.Ptr(true),
									KeepHTTPConnectionAlive:       to.Ptr(true),
									SupportHTTPStatusCodeContinue: to.Ptr(true),
									UnfoldHTTPHeaders:             to.Ptr(true),
								},
								EnvelopeSettings: &armlogic.AS2EnvelopeSettings{
									AutogenerateFileName:                    to.Ptr(true),
									FileNameTemplate:                        to.Ptr("<file-name-template>"),
									MessageContentType:                      to.Ptr("<message-content-type>"),
									SuspendMessageOnFileNameGenerationError: to.Ptr(true),
									TransmitFileNameInMimeHeader:            to.Ptr(true),
								},
								ErrorSettings: &armlogic.AS2ErrorSettings{
									ResendIfMDNNotReceived:  to.Ptr(true),
									SuspendDuplicateMessage: to.Ptr(true),
								},
								MdnSettings: &armlogic.AS2MdnSettings{
									DispositionNotificationTo:  to.Ptr("<disposition-notification-to>"),
									MdnText:                    to.Ptr("<mdn-text>"),
									MicHashingAlgorithm:        to.Ptr(armlogic.HashingAlgorithmSHA1),
									NeedMDN:                    to.Ptr(true),
									ReceiptDeliveryURL:         to.Ptr("<receipt-delivery-url>"),
									SendInboundMDNToMessageBox: to.Ptr(true),
									SendMDNAsynchronously:      to.Ptr(true),
									SignMDN:                    to.Ptr(true),
									SignOutboundMDNIfOptional:  to.Ptr(true),
								},
								MessageConnectionSettings: &armlogic.AS2MessageConnectionSettings{
									IgnoreCertificateNameMismatch: to.Ptr(true),
									KeepHTTPConnectionAlive:       to.Ptr(true),
									SupportHTTPStatusCodeContinue: to.Ptr(true),
									UnfoldHTTPHeaders:             to.Ptr(true),
								},
								SecuritySettings: &armlogic.AS2SecuritySettings{
									EnableNRRForInboundDecodedMessages:  to.Ptr(true),
									EnableNRRForInboundEncodedMessages:  to.Ptr(true),
									EnableNRRForInboundMDN:              to.Ptr(true),
									EnableNRRForOutboundDecodedMessages: to.Ptr(true),
									EnableNRRForOutboundEncodedMessages: to.Ptr(true),
									EnableNRRForOutboundMDN:             to.Ptr(true),
									OverrideGroupSigningCertificate:     to.Ptr(false),
								},
								ValidationSettings: &armlogic.AS2ValidationSettings{
									CheckCertificateRevocationListOnReceive: to.Ptr(true),
									CheckCertificateRevocationListOnSend:    to.Ptr(true),
									CheckDuplicateMessage:                   to.Ptr(true),
									CompressMessage:                         to.Ptr(true),
									EncryptMessage:                          to.Ptr(false),
									EncryptionAlgorithm:                     to.Ptr(armlogic.EncryptionAlgorithmAES128),
									InterchangeDuplicatesValidityDays:       to.Ptr[int32](100),
									OverrideMessageProperties:               to.Ptr(true),
									SignMessage:                             to.Ptr(false),
								},
							},
							ReceiverBusinessIdentity: &armlogic.BusinessIdentity{
								Qualifier: to.Ptr("<qualifier>"),
								Value:     to.Ptr("<value>"),
							},
							SenderBusinessIdentity: &armlogic.BusinessIdentity{
								Qualifier: to.Ptr("<qualifier>"),
								Value:     to.Ptr("<value>"),
							},
						},
						SendAgreement: &armlogic.AS2OneWayAgreement{
							ProtocolSettings: &armlogic.AS2ProtocolSettings{
								AcknowledgementConnectionSettings: &armlogic.AS2AcknowledgementConnectionSettings{
									IgnoreCertificateNameMismatch: to.Ptr(true),
									KeepHTTPConnectionAlive:       to.Ptr(true),
									SupportHTTPStatusCodeContinue: to.Ptr(true),
									UnfoldHTTPHeaders:             to.Ptr(true),
								},
								EnvelopeSettings: &armlogic.AS2EnvelopeSettings{
									AutogenerateFileName:                    to.Ptr(true),
									FileNameTemplate:                        to.Ptr("<file-name-template>"),
									MessageContentType:                      to.Ptr("<message-content-type>"),
									SuspendMessageOnFileNameGenerationError: to.Ptr(true),
									TransmitFileNameInMimeHeader:            to.Ptr(true),
								},
								ErrorSettings: &armlogic.AS2ErrorSettings{
									ResendIfMDNNotReceived:  to.Ptr(true),
									SuspendDuplicateMessage: to.Ptr(true),
								},
								MdnSettings: &armlogic.AS2MdnSettings{
									DispositionNotificationTo:  to.Ptr("<disposition-notification-to>"),
									MdnText:                    to.Ptr("<mdn-text>"),
									MicHashingAlgorithm:        to.Ptr(armlogic.HashingAlgorithmSHA1),
									NeedMDN:                    to.Ptr(true),
									ReceiptDeliveryURL:         to.Ptr("<receipt-delivery-url>"),
									SendInboundMDNToMessageBox: to.Ptr(true),
									SendMDNAsynchronously:      to.Ptr(true),
									SignMDN:                    to.Ptr(true),
									SignOutboundMDNIfOptional:  to.Ptr(true),
								},
								MessageConnectionSettings: &armlogic.AS2MessageConnectionSettings{
									IgnoreCertificateNameMismatch: to.Ptr(true),
									KeepHTTPConnectionAlive:       to.Ptr(true),
									SupportHTTPStatusCodeContinue: to.Ptr(true),
									UnfoldHTTPHeaders:             to.Ptr(true),
								},
								SecuritySettings: &armlogic.AS2SecuritySettings{
									EnableNRRForInboundDecodedMessages:  to.Ptr(true),
									EnableNRRForInboundEncodedMessages:  to.Ptr(true),
									EnableNRRForInboundMDN:              to.Ptr(true),
									EnableNRRForOutboundDecodedMessages: to.Ptr(true),
									EnableNRRForOutboundEncodedMessages: to.Ptr(true),
									EnableNRRForOutboundMDN:             to.Ptr(true),
									OverrideGroupSigningCertificate:     to.Ptr(false),
								},
								ValidationSettings: &armlogic.AS2ValidationSettings{
									CheckCertificateRevocationListOnReceive: to.Ptr(true),
									CheckCertificateRevocationListOnSend:    to.Ptr(true),
									CheckDuplicateMessage:                   to.Ptr(true),
									CompressMessage:                         to.Ptr(true),
									EncryptMessage:                          to.Ptr(false),
									EncryptionAlgorithm:                     to.Ptr(armlogic.EncryptionAlgorithmAES128),
									InterchangeDuplicatesValidityDays:       to.Ptr[int32](100),
									OverrideMessageProperties:               to.Ptr(true),
									SignMessage:                             to.Ptr(false),
								},
							},
							ReceiverBusinessIdentity: &armlogic.BusinessIdentity{
								Qualifier: to.Ptr("<qualifier>"),
								Value:     to.Ptr("<value>"),
							},
							SenderBusinessIdentity: &armlogic.BusinessIdentity{
								Qualifier: to.Ptr("<qualifier>"),
								Value:     to.Ptr("<value>"),
							},
						},
					},
				},
				GuestIdentity: &armlogic.BusinessIdentity{
					Qualifier: to.Ptr("<qualifier>"),
					Value:     to.Ptr("<value>"),
				},
				GuestPartner: to.Ptr("<guest-partner>"),
				HostIdentity: &armlogic.BusinessIdentity{
					Qualifier: to.Ptr("<qualifier>"),
					Value:     to.Ptr("<value>"),
				},
				HostPartner: to.Ptr("<host-partner>"),
				Metadata:    map[string]interface{}{},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
