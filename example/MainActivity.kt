package com.example.myapplication

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.view.View
import android.widget.Button
import android.widget.TextView
import android.widget.EditText

import test.Identity
import test.Test
import test.Global

class MainActivity : AppCompatActivity() {

    private var global = Global()
    private var _identity: Identity? = null

    private fun toHex(byteArray: ByteArray): String {
        val hexStringBuffer = StringBuffer()
        for (i in byteArray.indices) {
            hexStringBuffer.append(String.format("%02x", byteArray[i]))
        }
        return hexStringBuffer.toString()
    }

    private fun createIden() {
        val idTextView: TextView  = findViewById(R.id.id)
        idTextView.text = "Creating identity..."
        _identity = Test.newIdentity()
        val identity = _identity!!
        idTextView.text = String.format("id: \n %s\n", identity.id())
    }

    private fun sign() {
        val identity = _identity!!
        identity.unlock()
        val msgText: EditText = findViewById(R.id.msg)
        val msg = msgText.text.toString()
        val sig = identity.signKOp(msg.toByteArray())
        val sigTextView: TextView  = findViewById(R.id.signature)
        sigTextView.text = toHex(sig)
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val textView: TextView  = findViewById(R.id.id)
        textView.text = ""
        global.ksEthPath = String.format("%s/keyStoreEthereum", filesDir.absolutePath)
        global.pass = "password"
        Test.setGlobal(global)
        Test.initStorage()

        val createIdenButton: Button = findViewById(R.id.createIden)
        createIdenButton.setOnClickListener(object: View.OnClickListener {
            override fun onClick(v: View?) {
                createIden()
            }
        })
        val signButton: Button = findViewById(R.id.sign)
        signButton.setOnClickListener(object: View.OnClickListener {
            override fun onClick(v: View?) {
                sign()
            }
        })
    }
}
